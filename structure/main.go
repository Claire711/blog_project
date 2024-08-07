package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 定义一个结构体
type User struct {
	Name   string
	Gender string
	Age    int
}

func testpart(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Println("Prase template failed,err:%v", err)
		return
	}
	//渲染模板
	u1 := User{
		Name:   "Claire",
		Gender: "女",
		Age:    20,
	}
	t.Execute(w, u1)
}
func main() {
	http.HandleFunc("/", testpart)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start fail,err:%v", err)
		return
	}

}
