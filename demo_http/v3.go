package main

import (
	"log"
	"net/http"
	"time"
)

type myHandlerHttpV3 struct{}

func (*myHandlerHttpV3) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("this is version 3"))
}


func sayByeV3(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye ,this is v3 httpServer"))
}


func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandlerHttpV3{})
	mux.HandleFunc("/bye", sayByeV3)
	
	server := &http.Server{
		Addr:              ":1210",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 5,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	log.Println("Starting v3 httpserver")
	log.Fatal(server.ListenAndServe())
}