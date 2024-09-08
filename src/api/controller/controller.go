package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	ServeVideoController(c *fiber.Ctx) error
	VideoDetailsController(c *fiber.Ctx) error
}

type controller struct{}

func NewController() Controller {
	return &controller{}
}

func (c2 controller) ServeVideoController(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (c2 controller) VideoDetailsController(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
