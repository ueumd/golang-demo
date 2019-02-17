package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc  {
	return func(c *gin.Context) {
		reqId := c.Request.Header.Get("uuid")

		if reqId == "" {
			u4, _ := uuid.NewV4()
			reqId = u4.String()
		}
		c.Set("uuid", reqId)
		c.Writer.Header().Set("uuid", reqId)
		c.Next()
	}
}