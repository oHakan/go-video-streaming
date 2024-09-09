package main

import (
	"github.com/oHakan/go-video-streaming/helpers"
	"github.com/oHakan/go-video-streaming/src/api/controller"
	"github.com/oHakan/go-video-streaming/src/api/handler"
	"github.com/oHakan/go-video-streaming/src/internal/config"
	"github.com/oHakan/go-video-streaming/src/pkg/fiber"
	"log"
)

func main() {
	fiberAPI := fiber.NewFiberAPI()

	log.Print("Video streaming service has started.")
	config.InitializeConfig()

	port := config.GetPort()

	// Initialize the controller and handler
	mainStaticFolderDist := helpers.GetCurrentPath() + "/static"

	newController := controller.NewController(mainStaticFolderDist)
	newHandler := handler.NewHandler(newController)

	fiberAPI.Post("/upload-video", newHandler.UploadVideoHandler)
	fiberAPI.Static("/static", mainStaticFolderDist)
	fiberAPI.Get("/video-details", newHandler.VideoDetailsHandler)

	err := fiberAPI.Listen(":" + port)

	if err != nil {
		panic(err)
	}
}
