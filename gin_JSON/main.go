package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	/*	r.GET("/json", func(c *gin.Context) {*/
	//interface后是空接口，表示可以是任意类型的数据格式
	//data := map[string]interface{}{
	//	"name":    "claire",
	//	"message": "hello world",
	//	"age":     20,
	//	}
	/*	data := gin.H{
		"name":    "claire",
		"message": "hello world",
		"age":     20,
	}*/
	//方法2 结构体
	type msg struct {
		Name string
		Age  int
		Msg  string
		//灵活使用tag来做定制化操作
	}
	r.GET("/other", func(c *gin.Context) {
		data := msg{
			Name: "claire",
			Age:  20,
			Msg:  "hello world",
		}
		c.JSON(200, data)
	})

	r.Run(":9000")
}
