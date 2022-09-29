package main

import (
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetHttp(text string) string {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	bot, err := tgbotapi.NewBotAPI("5700919456:AAFdFEWCCOCRLu1rTcwFIrCgtJlg-u_4wno")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			text := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// bot.Send(msg)
		}
	}
}
