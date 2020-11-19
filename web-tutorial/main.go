package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	// [0]第一个文件
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		// ioutil.ReadAll把文件读取成[]byte
		data, err := ioutil.ReadAll(file)
		if err == nil {
			// string(data)把 []byte 字节切片转换为字符串，写如到 w responseWriter
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.PostFormValue("first_name"))
	})
	http.HandleFunc("/upload", upload)
	server.ListenAndServe()
}
