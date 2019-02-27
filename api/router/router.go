package router

import (
	"github.com/gin-gonic/gin"
	. "myapiserver/api/handler"
	"myapiserver/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine  {

	g.Use(mw...)
	
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 采用gin日志信息
	// r := gin.Default()

	g.POST("/login", Login)
	g.POST("/do", LoginTest)
	g.GET("/do", LoginGet)
	g.POST("welcome", LoginBind)
	g.POST("welcome2", LoginBind2)
	u := g.Group("/v1")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/getUser", GetUser)
		u.POST("/addUser", AddUser)
		u.PUT("/user/:id", Update)
		u.DELETE("/user/:id", Destroy)
	}

	return g
}