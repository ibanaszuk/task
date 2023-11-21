package main

import (
	"random-stuff-service/auth"
	"random-stuff-service/rest"
	"random-stuff-service/rest/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config := rest.Config{
		GinMode:       gin.ReleaseMode,
		Authenticator: auth.New(),
	}

	router := router.New(config)
	router.Run(":5000")
}
