package main

import "github.com/gin-gonic/gin"

func RunServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!!",
		})
	})
	r.GET("/app_version", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "Sample golang app version 1.0 !!",
                })
        })
	r.GET("/testing", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "New api added to the application !!",
                })
        })
	return r
}

func main() {
	RunServer().Run(":5000")
}

