package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/oHakan/go-video-streaming/docs"
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

// @title GoLang - Streaming Service
// @version 1.0
// @description This service provides your video CDN gateway
// @termsOfService http://swagger.io/terms/
// @contact.name oHakan
// @contact.email osmanhakan54@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func setupRoutes(app *fiber.App) {
	staticFolderPath := config.GetStaticFolderPath()
	mainStaticFolderDist := helpers.GetCurrentPath() + staticFolderPath

	ctrl := controller.NewController(mainStaticFolderDist)
	h := handler.NewHandler(ctrl)

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
	}))

	app.Post("/upload-video", h.UploadVideoHandler)
	app.Static("/static", mainStaticFolderDist)
	app.Get("/video-details", h.VideoDetailsHandler)
}
