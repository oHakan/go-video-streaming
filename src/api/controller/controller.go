package controller

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/oHakan/go-video-streaming/helpers"
	"log"
	"os/exec"
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

	fileNameWithoutExtension := strings.Split(file.Filename, ".")[0]
	isPathExists := helpers.IsDirectoryExists(c2.MainStaticFolderDist + "/" + fileNameWithoutExtension)

	if isPathExists {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	filePath := helpers.CreateNewStaticDirectory(c2.MainStaticFolderDist + "/" + fileNameWithoutExtension)

	createdFolder := filePath

	if createdFolder != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = c.SaveFile(file, fmt.Sprintf("%s/%s/%s", c2.MainStaticFolderDist, fileNameWithoutExtension, file.Filename))

	cmd := exec.Command("ffmpeg", "-i", fmt.Sprintf("%s/%s/%s", c2.MainStaticFolderDist, fileNameWithoutExtension, file.Filename), "-c:v", "libx264", "-preset", "veryfast", "-g", "48", "-keyint_min", "48", "-sc_threshold", "0", "-b:v", "2500k", "-maxrate", "2500k", "-bufsize", "5000k", "-c:a", "aac", "-b:a", "128k", "-hls_time", "10", "-hls_playlist_type", "vod", "-hls_segment_filename", fmt.Sprintf("%s/%s/output%%03d.ts", c2.MainStaticFolderDist, fileNameWithoutExtension), fmt.Sprintf("%s/%s/playlist.m3u8", c2.MainStaticFolderDist, fileNameWithoutExtension))

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Printf("ffmpeg error: %v: %s", err, stderr.String())
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("ffmpeg error: %v: %s", err, stderr.String()))
	}

	return c.SendStatus(fiber.StatusOK)
}

func (c2 controller) VideoDetailsController(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
