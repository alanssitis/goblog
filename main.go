package main

import (
	"io"
	"log"
	"net/http"
)

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	helloHandler := HelloHandler{}
	http.Handle("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
