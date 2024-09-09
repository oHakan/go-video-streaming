package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oHakan/go-video-streaming/src/api/controller"
)

type Handler interface {
	UploadVideoHandler(c *fiber.Ctx) error
	VideoDetailsHandler(c *fiber.Ctx) error
}

type handler struct {
	Controller controller.Controller
}

func NewHandler(controller controller.Controller) Handler {
	return &handler{Controller: controller}
}

func (h *handler) UploadVideoHandler(c *fiber.Ctx) error {
	if err := h.Controller.UploadVideoController(c); err != nil {
		return err // Fiber will handle the error response
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) VideoDetailsHandler(c *fiber.Ctx) error {
	if err := h.Controller.VideoDetailsController(c); err != nil {
		return err // Fiber will handle the error response
	}
	return c.SendStatus(fiber.StatusOK)
}
