package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberAPI() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:  true,
		EnableIPValidation: true,
	})

	// CORS Settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	return app
}
