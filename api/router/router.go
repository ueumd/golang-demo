package router

import "github.com/gin-gonic/gin"
import ."myapiserver/api/apis"

func InitRouter() *gin.Engine  {
	router := gin.Default()

	router.GET("/getUser", GetUser)

	router.POST("/addUser", AddUser)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	return router
}