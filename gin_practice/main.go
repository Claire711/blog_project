package main

import (
	"github.com/gin-gonic/gin"
)

//定义模板

func main() {
	r := gin.Default()

	//解析模板
	r.LoadHTMLFiles("./test.tmpl")
	//渲染模板
	//relativepath 后跟的是端口后面的网址内容
	r.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test.tmpl", "<a>claire的博客</a>")
	})
	//启动server
	r.Run(":8080")

}
