package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "myapiserver/api/handler"
	"myapiserver/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine  {
	cf := cors.DefaultConfig()  //新建cors配置
	cf.AllowAllOrigins = true   //允许跨域
	cf.AddAllowHeaders("token")


	g.Use(cors.New(cf))    //使用use方法让gin接受cors生成的配置

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