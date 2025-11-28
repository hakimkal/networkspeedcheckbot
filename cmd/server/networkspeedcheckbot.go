package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hakimkal/networkspeedcheckbot/internal/config"
	networkSpeedCheckService "github.com/hakimkal/networkspeedcheckbot/internal/service"
)

func main() {

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)

	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramApiToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("starting networkspeedcheckbot...")

	getTelegramBotMessages(bot)

}

func getTelegramBotMessages(bot *tgbotapi.BotAPI) {

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	var scs = networkSpeedCheckService.NewSpeedcheckService()

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {

		if update.Message == nil {
			continue
		}
		receivedMessage := update.Message
		var messageText string

		log.Printf("received update from telegram bot: %s", receivedMessage)

		if receivedMessage.Text == "/speed-check" {
			downloadSpeedMbps := scs.CheckDownloadSpeed()
			messageText = fmt.Sprintf("Download Speed: %.2f Mbps", downloadSpeedMbps)

		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)

		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}
