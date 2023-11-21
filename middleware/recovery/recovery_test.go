package recovery_test

import (
	"net/http"
	"net/http/httptest"
	"random-stuff-service/middleware/recovery"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecovery(t *testing.T) {
	router := gin.Default()
	router.Use(gin.CustomRecovery(recovery.ErrorHandler()))

	router.GET("/", func(ctx *gin.Context) {
		panic("recovery incoming")
	})

	request, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, 500, recorder.Code)
}
