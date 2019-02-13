package router

import (
	"github.com/gin-gonic/gin"
	. "myapiserver/api/apis"
	"myapiserver/middleware"
) //  .含义是这个包导入之后在你调用这个包的函数时，可以省略前缀的包名

func InitRouter() *gin.Engine  {
	router := gin.Default()

	router.GET("/getUser", GetUser)
	router.POST("/addUser", AddUser)
	router.PUT("/user/:id", Update)
	router.DELETE("/user/:id", Destroy)

	u := router.Group("/v1")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/login", GetAllUser)
	}

	return router
}