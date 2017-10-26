package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)

	log.Printf("Start on port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		linebot.WithHTTPClient(client),
	)

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				m := message.Text + "バジ〜"
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(m)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
