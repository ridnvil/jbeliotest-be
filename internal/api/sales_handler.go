package api

import (
	"github.com/gofiber/fiber/v2"
	"jubeliotesting/internal/service"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/middleware"
)

type SalesHandler struct {
	SalesService *service.SalesService
	Config       config.GetEnvConfig
}

func NewSalesHandler(salesService *service.SalesService, config config.GetEnvConfig) *SalesHandler {
	return &SalesHandler{SalesService: salesService, Config: config}
}

func (h *SalesHandler) SalesRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/sales/most-least", middleware.LoggingMiddleware(), middleware.APIKeyMiddleware(h.Config), h.MostLeastSales)
	api.Get("/sales/correlation", middleware.LoggingMiddleware(), middleware.APIKeyMiddleware(h.Config), h.CorrelationSalesApi)
	api.Get("/sales/bycountry", middleware.LoggingMiddleware(), middleware.APIKeyMiddleware(h.Config), h.CountrySalesApi)
}

func (h *SalesHandler) MostLeastSales(c *fiber.Ctx) error {
	data, err := h.SalesService.GetMostLeastSales()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

func (h *SalesHandler) CorrelationSalesApi(c *fiber.Ctx) error {
	groupBy := c.Query("groupBy")

	data, err := h.SalesService.GetCorrelationSales(groupBy)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

func (h *SalesHandler) CountrySalesApi(c *fiber.Ctx) error {
	data, err := h.SalesService.GetCountrySales()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}
