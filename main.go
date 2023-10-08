package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
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
	r.POST("/tg/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world "})
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			// Handle error
		}
		tgBawimySie(jsonData)
	})

	return r

}

func tgBawimySie(jsonData []byte) {

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		// print out if error is not nil
		fmt.Println(err)
	}
	message := ">> "
	for key, value := range result {
		val, ok := value.(string)
		if ok {
			message = message + "; " + key + ":" + val
			fmt.Println(key, ":", val)
			fmt.Println("=========")
			fmt.Println(message)
		}
		if !ok {
			fmt.Println(key)
			continue
		}
	}

	//message := result["result"].(string)

	url := "https://api.telegram.org/bot" + config.GetConf().BotId + "/sendMessage?chat_id=" + config.GetConf().ChatId + "&text=" + message

	jsonStr := []byte(`{}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:9981
	r.Run(":9981")
}
