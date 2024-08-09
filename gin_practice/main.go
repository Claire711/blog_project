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
	//如需加载静态文件，也应在解析前声明
	r.Static("/xyz", "./statics")
	//router.Static(relativePath, root)：
	//relativePath：请求的 URL 路径前缀。当用户访问 /xyz/ 路径时，Gin 会查找 root 目录中的文件。
	//root：静态文件的根目录。可以是项目中的相对路径或绝对路径。

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
