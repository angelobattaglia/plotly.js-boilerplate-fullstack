package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	}

func main() {
	http.HandleFunc("/public/index.html/")
	http.ListenAndServe(":8080", nil)
}