package controller

import (
	"fmt"
	"roybot/config"
	"roybot/service"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/gin-gonic/gin"
)

//Callback : Setup https server for receiving request from LINE platform
func Callback(c *gin.Context) {
	events, err := config.Bot.ParseRequest(c.Request)
	if err != nil {
		fmt.Printf("Request got error %+v", err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = config.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					service.CallAdmin("ReplyMessage got error", err)
				}
			}
		}
	}
}
