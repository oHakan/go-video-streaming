package main

import (
	"log"

	"github.com/oHakan/go-video-streaming/src/internal/config"
	"github.com/oHakan/go-video-streaming/src/internal/server"
)

func main() {
	config.InitializeConfig()

	app := server.SetupServer()

	port := config.GetPort()
	log.Printf("Video streaming service starting on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
