package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"myapiserver/pkg/errno"
)

type Response struct {
	Code		int			`json:"code"`
	Message 	string		`json:"message"`
	Data 		interface{}	`json:"data"`
}

/**
生成响应结果
 */
func SendResponse(c *gin.Context, err error, data interface{})  {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:	 code,
		Message: message,
		Data: 	 data,
	})
}