package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "I am GET method.",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "I am POST method.",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "I am PUT method.",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "I am DELETE method.",
		})
	})
	r.Run()
}
