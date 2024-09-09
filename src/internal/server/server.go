package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oHakan/go-video-streaming/src/api/controller"
	"github.com/oHakan/go-video-streaming/src/api/handler"
	"github.com/oHakan/go-video-streaming/src/internal/config"
	"github.com/oHakan/go-video-streaming/src/internal/helpers"
)

func SetupServer() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:  true,
		EnableIPValidation: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	setupRoutes(app)

	return app
}

func setupRoutes(app *fiber.App) {
	staticFolderPath := config.GetStaticFolderPath()
	mainStaticFolderDist := helpers.GetCurrentPath() + staticFolderPath

	ctrl := controller.NewController(mainStaticFolderDist)
	h := handler.NewHandler(ctrl)

	app.Post("/upload-video", h.UploadVideoHandler)
	app.Static("/static", mainStaticFolderDist)
	app.Get("/video-details", h.VideoDetailsHandler)
}
