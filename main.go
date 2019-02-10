package main

import (
	_ "myapiserver/api/database" // _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
	orm "myapiserver/api/database" // 别名
	"myapiserver/api/router"
)

func main()  {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}