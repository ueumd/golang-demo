package main

import (
	"flag"
	"fmt"
)

func main() {
	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "123", "name")

	// 当所有的flag声明完成后，使用`flag.Parse()`来分
	// 解命令行选项
	flag.Parse()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)

	/*
	ok: false
	id: 0
	port: :8080
	name: 123


	使用-h参数可以查看使用帮助：
	 go run flag.go -h
	  -id int
			id
	  -name string
			name (default "123")
	  -ok
			is ok
	  -port string
			http listen port (default ":8080")

	 */

}