package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/**
1. 首先关闭所有的监听;
2. 然后关闭所有的空闲连接;
3. 然后无限期等待连接处理完毕转为空闲，并关闭;
4. 如果提供了 带有超时的Context，将在服务关闭前返回 Context的超时错误;
*/

// 主动关闭服务器
var server *http.Server

type myHandlerV4 struct{}

func (*myHandlerV4) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

// 关闭http
func sayByeV4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server"))     // 没有输出

	err := server.Shutdown(nil)
	if err != nil {
		log.Println([]byte("shutdown the server err"))
	}
}

func main() {
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandlerV4{})
	mux.HandleFunc("/bye", sayByeV4)

	server = &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting v3 httpserver")
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")

}

// 尝试访问http://localhost:1210/bye 在控制台会得到以下提示结果，平滑关闭http服务成功:
