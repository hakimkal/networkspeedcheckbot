package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakimkal/networkspeedcheckbot/internal/keyboards"
	networkSpeedCheckService "github.com/hakimkal/networkspeedcheckbot/internal/service/speedcheck"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hello,You can run your server speed checks!."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SpeedCheck(bot *tgbotapi.BotAPI, update tgbotapi.Update, scs networkSpeedCheckService.SpeedcheckService) {
	latency := scs.PingTest()
	messageText := fmt.Sprintf("Latency: %.3f sec", latency)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func DownloadSpeedCheck(bot *tgbotapi.BotAPI, update tgbotapi.Update, scs networkSpeedCheckService.SpeedcheckService) {

	downloadSpeedMbps := scs.CheckDownloadSpeed()
	messageText := fmt.Sprintf("Download Speed: %.2f Mbps", downloadSpeedMbps)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func UploadSpeedCheck(bot *tgbotapi.BotAPI, update tgbotapi.Update, scs networkSpeedCheckService.SpeedcheckService) {
	uploadSpeedMbps := scs.CheckUploadSpeed()
	text := fmt.Sprintf("Upload Speed: %.2f Mbps\n", uploadSpeedMbps)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	 
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
