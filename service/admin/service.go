package admin

import (
	"fmt"
	"roybot/config"
	"time"

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

//Version return server information
func Version() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	result := fmt.Sprintf("ReleaseTime: %s\nReleaseVersion: %s\nSystemTime: %s",
		config.Conf.Release.Time,
		config.Conf.Release.Version,
		t)
	return result
}
