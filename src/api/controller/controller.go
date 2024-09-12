package controller

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/oHakan/go-video-streaming/src/internal/helpers"
)

const (
	allowedExtension = ".mp4"
)

type Controller interface {
	UploadVideoController(c *fiber.Ctx) error
	VideoDetailsController(c *fiber.Ctx) error
}

type controller struct {
	MainStaticFolderDist string
}

func NewController(mainStaticFolderDist string) Controller {
	return &controller{MainStaticFolderDist: mainStaticFolderDist}
}

// UploadVideoController
// ShowAccount godoc
// @Summary      Upload new video
// @Description  Upload new video to server. (Only .mp4 files allowed)
// @Tags         video
// @Accept       json
// @Produce      json
// @Param        file formData file true "Video file"
// @Success      200
// @Failure      500  Internal Server Error
// @Router       /upload-video [post]
func (ctrl *controller) UploadVideoController(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid file upload")
	}

	if filepath.Ext(file.Filename) != allowedExtension {
		return fiber.NewError(fiber.StatusBadRequest, "Only .mp4 files are allowed")
	}

	fileNameWithoutExt := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	fileDestination := filepath.Join(ctrl.MainStaticFolderDist, fileNameWithoutExt)

	if helpers.IsDirectoryExists(fileDestination) {
		return fiber.NewError(fiber.StatusConflict, "File already exists")
	}

	if err := helpers.CreateNewStaticDirectory(fileDestination); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process video")
	}

	uploadedFilePath := filepath.Join(fileDestination, file.Filename)
	if err := c.SaveFile(file, uploadedFilePath); err != nil {
		log.Printf("Failed to save file: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save video")
	}

	command := helpers.GenerateFFMPEGCommand(fileDestination, file.Filename)
	if err, stderr := helpers.RunCommand(command); err != nil {
		log.Printf("ffmpeg error: %v: %s", err, stderr.String())
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process video")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *controller) VideoDetailsController(c *fiber.Ctx) error {
	// Implement video details logic here
	return c.SendStatus(fiber.StatusNotImplemented)
}
