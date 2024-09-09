package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/oHakan/go-video-streaming/helpers"
	"log"
	"strings"
)

type Controller interface {
	UploadVideoController(c *fiber.Ctx) error
	VideoDetailsController(c *fiber.Ctx) error
}

type controller struct {
	MainStaticFolderDist string
}

func NewController(mainStaticFolderDist string) Controller {
	return &controller{
		mainStaticFolderDist,
	}
}

func (c2 controller) UploadVideoController(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	splittedFileName := strings.Split(file.Filename, ".")

	if splittedFileName[1] != "mp4" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fileNameWithoutExtension := splittedFileName[0]
	fileDestination := c2.MainStaticFolderDist + "/" + fileNameWithoutExtension
	isPathExists := helpers.IsDirectoryExists(fileDestination)

	if isPathExists {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	createdFilePath := helpers.CreateNewStaticDirectory(fileDestination)

	if createdFilePath != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	uploadedFilePath := fmt.Sprintf("%s/%s", fileDestination, file.Filename)
	err = c.SaveFile(file, uploadedFilePath)

	command := helpers.GenerateFFMPEGCommand(fileDestination, file.Filename)

	err, stderr := helpers.RunCommand(command)

	if err != nil {
		log.Printf("ffmpeg error: %v: %s", err, stderr.String())
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("ffmpeg error: %v: %s", err, stderr.String()))
	}

	return c.SendStatus(fiber.StatusOK)
}

func (c2 controller) VideoDetailsController(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
