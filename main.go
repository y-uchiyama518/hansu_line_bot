package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// create line bt client
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// create text message
	message := linebot.NewTextMessage("hello world")
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
