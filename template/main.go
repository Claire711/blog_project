package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func testPart(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err:%v", err)
		return
	}

	//渲染
	err = t.Execute(w, "小公主")
	if err != nil {
		fmt.Println("rander template failed,err:%v", err)
		return
	}

}
func main() {
	http.HandleFunc("/", testPart)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		println("ListenAndServe can't start server")
		return
	}

}
