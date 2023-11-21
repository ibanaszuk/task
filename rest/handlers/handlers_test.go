package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"random-stuff-service/rest"
	"random-stuff-service/rest/router"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testSuite struct {
	t      *testing.T
	router *gin.Engine
	// TODO: send in auth http header once auth.ValidateAccessToken() is being called in the API handler
	// accessToken string
}

func createTestSuite(t *testing.T) testSuite {
	//TODO: use once auth.ValidateAccessToken() is being called in the API handler
	// authenticator, err := auth.New()
	// assert.NoError(t, err)

	config := rest.Config{
		GinMode: "debug",
		Client:  http.DefaultClient,
	}

	router := router.New(config)
	go func() {
		err := http.ListenAndServe(":5000", router)
		assert.NoError(t, err)
	}()

	return testSuite{
		t:      t,
		router: router,
	}
}

func TestHandlers(t *testing.T) {
	testSuite := createTestSuite(t)
	testSuite.runGetTests()
}

func (ts testSuite) runGetTests() {
	ts.t.Run("GetRandomNameAndJoke", func(t *testing.T) {
		url := "api/v1/random-name-with-joke"
		request, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		// TODO: send in auth http header once auth.ValidateAccessToken() is being called in the API handler
		//request.Header.Add("Authorization", "Bearer " + ts.accessToken)

		recorder := httptest.NewRecorder()
		ts.router.ServeHTTP(recorder, request)
		assert.Equal(t, 200, recorder.Code)
		assert.NotEmpty(t, recorder.Body)
	})
}
