package rest

import "net/http"

type Config struct {
	GinMode string
	Client  *http.Client
	//Authenticator authenticator.Authenticator
}
