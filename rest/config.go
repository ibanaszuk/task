package rest

import (
	"net/http"
	"random-stuff-service/auth"
)

type Config struct {
	GinMode       string
	Client        *http.Client
	Authenticator auth.Auth
}
