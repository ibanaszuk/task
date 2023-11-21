package ratelimiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Message struct {
	Body string `json:"body"`
}

func Ratelimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(10, 25)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(429, Message{
				Body: "API exceeded rate limit",
			})
		}
	}
}
