package main

import (
	"net/http"
	"io"
	"log"
)

type myHandler struct {}

/**
Golang Http Server源码阅读
http://www.cnblogs.com/yjf512/archive/2012/08/22/2650873.html

myHandler 实现了 net/http/server.go 接口中的 ServeHTTP方法

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}


 */
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.String()
	switch path {
	case "/":
		io.WriteString(w, "<h1>root</h1><a href=\"hi\">hi</a>")
	case "/hi":
		io.WriteString(w, "<h1>abc</h1><a href=\"/\">root</a>")
	}
}


func main()  {
	err := http.ListenAndServe(":5000", &myHandler{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}