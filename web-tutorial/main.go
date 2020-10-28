package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {
	httDir := http.Dir("D:\\Data\\from_github\\learning_go\\web-tutorial\\www")
	fmt.Println(httDir)
	http.ListenAndServe("localhost:8080", http.FileServer(httDir))
}
