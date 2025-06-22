package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"jubeliotesting/internal/dto"
	"jubeliotesting/internal/service"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/middleware"
	"log"
)

type PublishHandler struct {
	PublishService *service.PublisherService
	Config         config.GetEnvConfig
}

func NewPublishHandler(publishService *service.PublisherService, config config.GetEnvConfig) *PublishHandler {
	return &PublishHandler{PublishService: publishService, Config: config}
}

func (h *PublishHandler) PublishRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/publish", middleware.LoggingMiddleware(), middleware.APIKeyMiddleware(h.Config), h.Publish)
}

func (h *PublishHandler) Publish(c *fiber.Ctx) error {

	var dataBody struct {
		ClientID   string `json:"client_id"`
		MasterData bool   `json:"master_data_only"`
	}
	if err := c.BodyParser(&dataBody); err != nil {
		log.Println(err)
	}

	var pubilshDto dto.PublishDto
	pubilshDto.MasterData = dataBody.MasterData
	pubilshDto.Key = c.Get("X-API-KEY")
	pubilshDto.ClientID = dataBody.ClientID

	if err := h.PublishService.PublishMessage(context.Background(), pubilshDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{"status": "ok", "message": "process data running"})
}
