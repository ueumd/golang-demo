package router

import (
	"github.com/gin-gonic/gin"
	. "myapiserver/api/apis"
	"myapiserver/middleware"
)

func InitRouter() *gin.Engine  {
	router := gin.Default()

	router.POST("/login", Login)

	u := router.Group("/v1")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/getUser", GetUser)
		u.POST("/addUser", AddUser)
		u.PUT("/user/:id", Update)
		u.DELETE("/user/:id", Destroy)
	}

	return router
}