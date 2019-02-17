package middleware

import "github.com/gin-gonic/gin"
import (
	"myapiserver/pkg/token"
	. "myapiserver/pkg/result"
	"myapiserver/pkg/errno"
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