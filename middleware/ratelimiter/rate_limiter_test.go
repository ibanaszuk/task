package ratelimiter_test

import (
	"net/http"
	"net/http/httptest"
	"random-stuff-service/middleware/ratelimiter"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	router := gin.Default()
	router.Use(ratelimiter.Ratelimiter())

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "")
	})

	request, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	maxRequests := 25
	for i := 0; i < 100; i++ {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		if i > maxRequests {
			assert.Equal(t, 429, recorder.Code)
		}
	}
}
