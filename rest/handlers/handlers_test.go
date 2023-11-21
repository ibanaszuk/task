package handlers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type testSuite struct {
	t           *testing.T
	router      *gin.Engine
	accessToken string
}

func createTestSuite(t *testing.T) {
	//TODO: use once auth.ValidateAccessToken() is being called in the API handler
	// authenticator, err := auth.New()
	// assert.NoError(t, err)

}
