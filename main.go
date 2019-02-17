package main

import (
	orm "myapiserver/api/database" // 别名
	"myapiserver/api/router"
		"github.com/gin-gonic/gin"
	"myapiserver/middleware"
		"myapiserver/config"
)

func main()  {

	config.InitConfig()

	// database init
	orm.Init()
	defer orm.Eloquent.Close()

	// gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	router := router.Load(
		g,
		middleware.Logging(),
		middleware.RequestId(),
	)
	router.Run(":8000")
}