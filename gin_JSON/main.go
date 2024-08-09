package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//interface后是空接口，表示可以是任意类型的数据格式
		data := map[string]interface{}{
			"name":    "claire",
			"message": "hello world",
			"age":     20,
		}
		c.JSON(200, data)
	})
	r.Run(":9000")
}
