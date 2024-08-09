package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		//获取query string类型的参数
		//定义name变量携带query string类型的参数git
		/*name := c.Query("query")*/
		//取不到就用默认值
		/*name := c.DefaultQuery("query", "who")*/
		name, ok := c.GetQuery("query")
		if !ok {
			//返回一个默认值
			//赋值用=
			name =
				"somebody"
		}
		c.JSON(200, gin.H{
			//返回name所附带的参数
			//键值对用：git
			"name": name,
		})
	})

	//让server服务器运行
	r.Run(":8080")
}
