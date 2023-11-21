package auth_test

import (
	"fmt"
	"random-stuff-service/auth"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateAccessToken(t *testing.T) {
	auth := auth.New()

	var ctx *gin.Context
	accessToken := "ewogICAgInBlcm1pc3Npb25zIjogWwogICAgICAgICJQT1NUIiwKICAgICAgICAiR0VUIgogICAgXQp9"
	authHeader := fmt.Sprintf("Bearer %s", accessToken)
	ctx.Header("authorization", authHeader)

	err := auth.ValidateAccessToken(ctx, "POST")
	assert.NoError(t, err)

	err = auth.ValidateAccessToken(ctx, "PUT")
	assert.Error(t, err)
}
