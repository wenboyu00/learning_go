package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := loadTemplates()
	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			// 根据请求的路径，找到响应的模板名
			fileName := request.URL.Path[1:]
			t := templates.Lookup(fileName)
			if t != nil {
				err := t.Execute(writer, nil)
				if err != nil {
					log.Fatalln(err.Error())
				}
			} else {
				writer.WriteHeader(http.StatusNotFound)
			}
		})
	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot")))
	http.ListenAndServe("localhost:8080", nil)
}

func loadTemplates() *template.Template {
	// 建立模板容器
	result := template.New("templates")
	// 解析出来的模板名，就是文件名
	result, err := result.ParseGlob("templates/*.html")
	// Must检查建模是否解析错误
	template.Must(result, err)
	return result
}
