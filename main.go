package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// response headers
	Header http.Header
	// response body
	Body io.ReadCloser
	// request that was sent to obtain the response
	Request *http.Request
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ANSWER": "PONG"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})
	r.POST("/tg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})
	r.POST("/tg/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
	})

	return r

}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:9981
	r.Run(":9981")
}
