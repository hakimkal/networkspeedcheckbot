package service

import (
	"fmt"
	"log"

	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedcheckService struct {
	Server *speedtest.Server
}

//Constructor

func NewSpeedcheckService() *SpeedcheckService {
	server := initializeServer()
	return &SpeedcheckService{Server: server}
}
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

func (scs *SpeedcheckService) PingTest() {

	server := scs.Server
	// Perform ping test (latency)
	var err = server.PingTest(nil)
	if err != nil {
		log.Fatalf("Error during ping test: %v", err)
	}
	fmt.Printf("Latency: %s\n", server.Latency)

}

func (speedCheckService *SpeedcheckService) CheckDownloadSpeed() float64 {
	// Perform download test
	server := speedCheckService.Server
	var err = server.DownloadTest()
	if err != nil {
		log.Fatalf("Error during download test: %v", err)
	}
	// Convert DLSpeed from bits per second to Mbps
	downloadSpeedMbps := float64(server.DLSpeed) / 1_000_000 * 8
	return downloadSpeedMbps
}

func (speedcheckService *SpeedcheckService) CheckUploadSpeed() {
	server := speedcheckService.Server
	// Perform upload test
	err := server.UploadTest()
	if err != nil {
		log.Fatalf("Error during upload test: %v", err)
	}
	// Convert ULSpeed from bits per second to Mbps
	uploadSpeedMbps := float64(server.ULSpeed) / 1_000_000 * 8
	fmt.Printf("Upload Speed: %.2f Mbps\n", uploadSpeedMbps)
}
