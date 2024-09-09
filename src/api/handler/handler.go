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
	return &handler{
		Controller: controller,
	}
}

func (h handler) VideoDetailsHandler(c *fiber.Ctx) error {
	return h.Controller.VideoDetailsController(c)
}

func (h handler) UploadVideoHandler(c *fiber.Ctx) error { return h.Controller.UploadVideoController(c) }
