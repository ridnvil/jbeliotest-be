package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type UploadDataSetHandler struct{}

func NewUploadDataSetHandler() *UploadDataSetHandler {
	return &UploadDataSetHandler{}
}

func (h *UploadDataSetHandler) UploadDataset(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dst := fmt.Sprintf("./uploads/%s", fileHeader.Filename)
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
