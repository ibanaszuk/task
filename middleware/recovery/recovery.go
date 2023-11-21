package recovery

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Description interface{} `json:"description"`
}

func ErrorHandler() gin.RecoveryFunc {
	return func(ctx *gin.Context, err any) {
		response := Response{Description: err}
		ctx.AbortWithStatusJSON(500, response)
	}
}
