package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/middleware"
	"path"
)

type UploadDataSetHandler struct {
	Config config.GetEnvConfig
}

func NewUploadDataSetHandler(conf config.GetEnvConfig) *UploadDataSetHandler {
	return &UploadDataSetHandler{Config: conf}
}

func (h *UploadDataSetHandler) Route(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/uploads", middleware.LoggingMiddleware(), middleware.APIKeyMiddleware(h.Config), h.UploadDataset)
}

func (h *UploadDataSetHandler) UploadDataset(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("files")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get extention
	ext := path.Ext(fileHeader.Filename)

	if ext != ".xlsx" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors.New("file type not supported"),
		})
	}

	fileName := "dataset"
	dst := fmt.Sprintf("./dataset/%s%s", fileName, ext)
	if err := c.SaveFile(fileHeader, dst); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    dst,
		"message": "success",
	})
}
