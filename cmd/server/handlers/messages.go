package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakimkal/networkspeedcheckbot/internal/service"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	service.SetTaskCallback(bot, update)
}
