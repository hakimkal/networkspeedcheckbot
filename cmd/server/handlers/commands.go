package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakimkal/networkspeedcheckbot/internal/service"
	networkSpeedCheckService "github.com/hakimkal/networkspeedcheckbot/internal/service/speedcheck"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var scs = networkSpeedCheckService.NewSpeedcheckService()
	log.Printf("Received Command...%s", update.Message.Command())
	switch update.Message.Command() {
	case "start":
		service.Start(bot, update)
	case "speedCheck":
		service.SpeedCheck(bot, update, *scs)
	case "downloadSpeedCheck":
		service.DownloadSpeedCheck(bot, update, *scs)

	case "uploadSpeedCheck":
		service.UploadSpeedCheck(bot, update, *scs)

	}
}
