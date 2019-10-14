package main

import (
	"log"
	"net/http"
)

func sayBye(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("bye bye, this is v1 httpServer"))
}

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)

	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":1210", nil))
}

