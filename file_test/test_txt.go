package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
func main() {
	r := gin.Default()
	r.GET("/", sayHello)
	//启动服务
	r.Run()
}
