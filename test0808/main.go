package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func praise(w http.ResponseWriter, r *http.Request) {
	kua := func(name string) (string, error) {
		return "帅气的" + name, nil
	}

	//解析模板

	//这里本来想直接用template.parse,但是好像要把它拆开才可以
	//定义模板
	//template是模板的意思，我这里的模板不只是要打开这个文件，还有自定义的函数，所以不能直接用template.parse
	t := template.New("test.tmpl")

	//告诉模板引擎，我现在多了一个自定义函数kua
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	_, err := t.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:%v", err)
		return
	}
	name := "claire"
	//渲染模板

	t.Execute(w, name)

}

func main() {
	http.HandleFunc("/", praise)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed,err:%v", err)
		return
	}
}
