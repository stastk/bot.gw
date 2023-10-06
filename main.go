package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})
	r.POST("/tg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})

	r.Run(":9981")
}
