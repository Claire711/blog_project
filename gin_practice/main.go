package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

//定义模板

func main() {
	r := gin.Default()

	//在解析模板之前，告诉电脑，我还有自定义的部分
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		}, //这里需要在后面加上，
	})
	//解析模板
	r.LoadHTMLFiles("./test.tmpl")
	//渲染模板
	//relativepath 后跟的是端口后面的网址内容
	r.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
	//启动server
	r.Run(":8080")

}
