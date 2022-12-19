package middleware

import (
	"github.com/NetSinx/go-restful-API/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		getAuthor := ctx.Request.Header.Get("Authorization")

		if getAuthor == "" {
			utils.Unauthorized(ctx)
			ctx.Abort()

			return
		}

		err := ValidateToken(getAuthor); if err != nil {
			utils.Unauthorized(ctx)
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}