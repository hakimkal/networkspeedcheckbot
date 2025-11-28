package client

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(TELEGRAM_API_KEY string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_API_KEY)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
