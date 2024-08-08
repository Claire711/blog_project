package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func good(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.New("test.tmpl").Delims("{[", "]}").ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
		return
	}
	name := "girl"
	//渲染模板
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", good)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed,err:%v", err)
		return
	}
}
