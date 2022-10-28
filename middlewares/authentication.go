package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rasyidev/dts-h8-mygram/helper"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
