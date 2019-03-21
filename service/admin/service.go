package admin

import (
	"fmt"
	"roybot/config"

	"github.com/line/line-bot-sdk-go/linebot"
)

// CallAdmin is a func for Admin..
func CallAdmin(method string, obj interface{}) {
	str := method
	if obj != nil {
		str += fmt.Sprintf(" : %+v ", obj)
	}
	_, err := config.Bot.PushMessage(config.Conf.Linebot.AdminID, linebot.NewTextMessage(str)).Do()
	if err != nil {
		fmt.Printf("Linebot Send message Get err.\n%+v ", err)
	}
}
