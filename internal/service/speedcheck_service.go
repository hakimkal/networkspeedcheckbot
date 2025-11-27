package service

import (
	"fmt"
	"log"

	"github.com/showwin/speedtest-go/speedtest"
)

func initializeServer() *speedtest.Server {

	// Fetch a list of available speed test servers
	serverList, err := speedtest.FetchServers()
	if err != nil {
		log.Fatalf("Error fetching server list: %v", err)
	}

	// Find the closest server(s). You can specify a list of server IDs if needed.
	// For simplicity, we'll just take the first closest server.
	closestServers, err := serverList.FindServer([]int{}) // Empty slice finds closest
	if err != nil {
		log.Fatalf("Error finding target servers: %v", err)
	}

	server := closestServers[0]
	return server

}
func PingTest() {
	server := initializeServer()
	// Perform ping test (latency)
	var err = server.PingTest(nil)
	if err != nil {
		log.Fatalf("Error during ping test: %v", err)
	}
	fmt.Printf("Latency: %s\n", server.Latency)

}

func CheckDownloadSpeed() {

}

func CheckUploadSpeed() {

}
