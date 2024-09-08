package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberAPI() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:  true,
		EnableIPValidation: true,
	})

	return app
}
