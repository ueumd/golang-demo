package middleware

import "github.com/gin-gonic/gin"
import (
	"myapiserver/pkg/token"
	"myapiserver/pkg/core"
	"myapiserver/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		if _, err := token.ParseRequest(ctx); err != nil {
			core.SendResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}