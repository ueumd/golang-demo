package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			fmt.Fprintf(w, "Hello333 %s\n", req.FormValue("name"))
			return
		}
		http.NotFound(w, req)
	})

	http.HandleFunc("/form", MyForm)

	log.Fatalf("%v", http.ListenAndServe("localhost:8081", nil))
}

func MyForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	formData := r.Form
	log.Printf("收到的数据： %v", formData)

	fmt.Fprintf(w, "提交成功")
}