package middleware

import "github.com/gin-gonic/gin"
import (
	"myapiserver/pkg/errno"
	"myapiserver/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		if _, err := token.ParseRequest(ctx); err != nil {
			SendResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}