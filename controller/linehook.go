package controller

import (
	"fmt"
	"roybot/config"
	"roybot/service/admin"
	"roybot/service/chat"

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
				s := chat.ParseMessage(message.Text)
				if len(s) > 0 {
					if _, err = config.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(s)).Do(); err != nil {
						admin.CallAdmin("ReplyMessage got error", err)
					}
				}

			case *linebot.StickerMessage:
				if _, err = config.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(chat.GetRandomReplay())).Do(); err != nil {
					admin.CallAdmin("ReplyMessage got error", err)
				}
			}
		}
	}
}

func TestCallback(c *gin.Context) {
	admin.CallAdmin(chat.ParseMessage("rp fs"), nil)
}
