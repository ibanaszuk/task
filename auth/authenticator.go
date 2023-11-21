package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	// Domain       string
	// Audience     string
	// ClientId     string
	// ClientSecret string
}

type Claims struct {
	Permissions []string `json:"permissions"`
}

func New() Auth {
	return Auth{
		//TODO: once auth0 tenant is set up, set these env vars between envs
		// Domain:       os.Getenv("AUTH0_DOMAIN"),
		// Audience:     os.Getenv("AUTH0_AUDIENCE"),
		// ClientId:     os.Getenv("AUTH0_CLIENT_ID"),
		// ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
	}
}

func (a Auth) ValidateAccessToken(ctx *gin.Context, permission string) error {
	encodedClaims, err := getEncodedClaims(ctx)
	if err != nil {
		return err
	}

	claims, err := newClaims(encodedClaims)
	if err != nil {
		return err
	}

	err = claims.validatePermissions(permission)
	if err != nil {
		return err
	}

	return nil
}

func getEncodedClaims(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get("authorization")
	if bearerToken == "" {
		return "", errors.New("auth header is missing")
	}

	token := strings.ReplaceAll(bearerToken, "Bearer ", "")
	parsedToken := strings.Split(token, ".")
	if len(parsedToken) < 2 {
		return "", errors.New("the following token is invalid: " + token)
	}

	encodedClaims := parsedToken[1]
	if encodedClaims == "" {
		return "", errors.New("the following token has empty claims: " + token)
	}

	return encodedClaims, nil
}

func newClaims(encodedClaims string) (Claims, error) {
	decodedClaims, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(encodedClaims)
	if err != nil {
		return Claims{}, err
	}

	var tokenClaims Claims
	err = json.Unmarshal(decodedClaims, &tokenClaims)
	if err != nil {
		return Claims{}, err
	}

	return tokenClaims, nil
}

func (c Claims) validatePermissions(permission string) error {
	if permission == "" {
		return errors.New("token validator received empty permission")
	}

	hasValidPermission := doesExist(c.Permissions, permission)
	if !hasValidPermission {
		return errors.New("invalid permissions")
	}

	return nil
}

func doesExist(permissions []string, incomingPermission string) bool {
	for _, permission := range permissions {
		if incomingPermission == permission {
			return true
		}
	}

	return false
}
