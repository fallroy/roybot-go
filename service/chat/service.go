package chat

import (
	"fmt"
	"math/rand"
	"roybot/service/finance"
	"roybot/service/lottery"
	"time"
)

// ParseMessage is a func
func ParseMessage(m string) string {
	result := ""
	fmt.Println(m)
	if m == "rp ft" {
		result = finance.GetFuture()
	} else if m == "rp fs" {
		result = finance.GetForeignFuture()
	} else if m == "lt s" {
		result = lottery.GetSummery()
	} else {
		result = GetRandomReplay()
	}
	return result
}

// GetRandomReplay is a func
func GetRandomReplay() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(7)
	return replayMessage[i]
}

var replayMessage = map[int]string{
	0: "Did you think...? I am just a BOT",
	1: "My dear guest. I can't answer.",
	2: "To be or not to be, this is a question.",
	3: "I LOVE YOU!",
	4: "What!?",
	5: "My name is RoyBot, I come from Taiwan.",
	6: "Seeing you make me angry!",
	7: "You! shut up!",
}
