package main

import (
	"log"
	"net/http"
)

type myHandlerHttp struct{}


func (*myHandlerHttp) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("this is version 2"))
}


func sayByeV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v2 httpServer"))
}


func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandlerHttp{})

	mux.HandleFunc("/bye", sayByeV2)

	log.Println("Starting v2 httpserver")
	log.Fatal(http.ListenAndServe(":1210", mux))
}


