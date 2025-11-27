package main

import (
	"log"

	"github.com/hakimkal/networkspeedcheckbot/internal/config"
	"github.com/hakimkal/networkspeedcheckbot/internal/service"
)

func main() {

	// Load config
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)

	}

	log.Printf("starting networkspeedcheckbot...")

	service.PingTest()
}
