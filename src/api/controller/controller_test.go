package controller_test

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/oHakan/go-video-streaming/src/api/controller"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadVideoController_Success(t *testing.T) {
	app := fiber.New()
	ctrl := controller.NewController("/tmp")
	app.Post("/upload-video", ctrl.UploadVideoController)
	os.RemoveAll("/tmp/test-video")

	file, err := os.Open("test_data/test-video.mp4")
	if err != nil {
		t.Fatalf("Failed to open test video file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/upload-video", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 3*1000)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	assert.Contains(t, string(responseBody), "http://localhost:9000")
}

func TestUploadVideoController_InvalidFileUpload(t *testing.T) {
	app := fiber.New()
	ctrl := controller.NewController("/tmp")
	app.Post("/upload-video", ctrl.UploadVideoController)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload-video", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 3*1000)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUploadVideoController_InvalidFileExtension(t *testing.T) {
	app := fiber.New()
	ctrl := controller.NewController("/tmp")
	app.Post("/upload-video", ctrl.UploadVideoController)

	file, err := os.Open("test_data/test-video.txt")
	if err != nil {
		t.Fatalf("Failed to open test video file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/upload-video", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 3*1000)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUploadVideoController_FileAlreadyExists(t *testing.T) {
	app := fiber.New()
	ctrl := controller.NewController("/tmp")
	app.Post("/upload-video", ctrl.UploadVideoController)

	_ = os.MkdirAll("/tmp/test-video", os.ModePerm)
	defer os.RemoveAll("/tmp/test-video")

	file, err := os.Open("test_data/test-video.mp4")
	if err != nil {
		t.Fatalf("Failed to open test video file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/upload-video", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 3*1000)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	assert.Equal(t, fiber.StatusConflict, resp.StatusCode)
}

func TestUploadVideoController_FailedToSaveFile(t *testing.T) {
	app := fiber.New()
	ctrl := controller.NewController("/tmp")
	app.Post("/upload-video", ctrl.UploadVideoController)

	file := bytes.NewReader([]byte("invalid content"))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "test-video.mp4")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/upload-video", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req, 3*1000)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}
