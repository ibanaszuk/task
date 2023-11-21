package router

import (
	"random-stuff-service/middleware/ratelimiter"
	"random-stuff-service/rest"
	"random-stuff-service/rest/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(config rest.Config) *gin.Engine {
	gin.SetMode(config.GinMode)
	router := gin.Default()
	router.Use(cors.Default()) //TODO: replace with custom cors middleware
	router.Use(ratelimiter.Ratelimiter())
	//router.Use()

	//TODO: limitTo, firstName, and lastName should be optional query params
	router.GET("/random-name-with-joke", handlers.Get(config))
	return router
}
