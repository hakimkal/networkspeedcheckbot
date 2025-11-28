package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakimkal/networkspeedcheckbot/internal/service"
	networkSpeedCheckService "github.com/hakimkal/networkspeedcheckbot/internal/service/speedcheck"
	"github.com/hakimkal/networkspeedcheckbot/internal/util"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, _ := util.GetKeyValue(update.CallbackQuery.Data)
	 
	var scs = networkSpeedCheckService.NewSpeedcheckService()

	switch {
	case cmd == "start":
		service.Start(bot, update)
	case cmd == "speedCheck":
		service.SpeedCheck(bot, update, *scs)

	case cmd == "downloadSpeedCheck":
		service.DownloadSpeedCheck(bot, update, *scs)

	case cmd == "uploadSpeedCheck":
		service.UploadSpeedCheck(bot, update, *scs)

	}

}
