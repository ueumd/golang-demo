package main

import (
	_ "myapiserver/api/database"
	orm "myapiserver/api/database"
	"myapiserver/api/router"
)

func main()  {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}