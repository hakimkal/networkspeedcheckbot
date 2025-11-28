package main

import (
	"log"

	"github.com/hakimkal/networkspeedcheckbot/cmd/client"
	"github.com/hakimkal/networkspeedcheckbot/cmd/server/handlers"
	"github.com/hakimkal/networkspeedcheckbot/internal/config"
)

func main() {

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)

	}

	bot := client.Init(cfg.TelegramApiToken)
	handlers.Init(bot)

	log.Printf("starting networkspeedcheckbot...")

}
