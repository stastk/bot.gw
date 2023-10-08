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

// j := []byte(`{message :
// 	map[
// 		chat:map[
// 				first_name:Stas
// 				id:2.32948736e+08
// 				type:private
//  			username:UXtas
// 			]
// 		date:1.696748623e+09
// 		from:map[
// 				first_name:Stas
// 				id:2.32948736e+08
// 				is_bot:false
// 				language_code:ru
// 				username:UXtas
// 		]
// 		message_id:136
// 		text:123
// 	]}`

type Chat struct {
	FirstName string `json:"first_name"`
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Username  string `json:"username"`
}

type From struct {
	FirstName    string `json:"first_name"`
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	Username     string `json:"username"`
}

type Message struct {
	Chat      Chat   `json:"chat"`
	Date      int    `json:"date"`
	From      From   `json:"from"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
}

type Answer struct {
	Message  Message `json:"message"`
	UpdateId string  `json:"update_id"`
}

func tgBawimySie(jsonData []byte) {
	var answer Answer
	//Data := []byte(jsonData)
	err := json.Unmarshal(jsonData, &answer)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Answer:", answer)
	fmt.Println("Answer message:", answer.Message)
	fmt.Println("Name is:", answer.UpdateId)

	url := "https://api.telegram.org/bot" + config.GetConf().BotId + "/sendMessage?chat_id=" + config.GetConf().ChatId + "&text=" + answer.UpdateId

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
