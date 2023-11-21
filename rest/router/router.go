package router

import (
	"random-stuff-service/rest"
	"random-stuff-service/rest/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(config rest.Config) *gin.Engine {
	gin.SetMode(config.GinMode)
	router := gin.Default()
	router.Use(cors.Default()) //replace with custom cors middleware
	//router.Use() custom rate limiter middleware
	//router.Use() custom recovery middleware

	group := router.Group("/api/v1")
	group.GET("/random-name-with-joke", handlers.Get(config))
	return router
}
