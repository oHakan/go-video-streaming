package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oHakan/go-video-streaming/helpers"
	"github.com/oHakan/go-video-streaming/src/api/controller"
	"github.com/oHakan/go-video-streaming/src/api/handler"
	"github.com/oHakan/go-video-streaming/src/internal/config"
	"log"
)

func main() {
	fiberAPI := fiber.New()

	fiberAPI.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:63342",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	log.Print("Video streaming service has started.")
	config.InitializeConfig()

	port := config.GetPort()

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
