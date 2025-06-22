package middleware

import (
	"github.com/gofiber/fiber/v2"
	"jubeliotesting/pkg/config"
)

func APIKeyMiddleware(config config.GetEnvConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-KEY")

		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": fiber.StatusUnauthorized,
				"error":  "API key is required",
			})
		}

		if apiKey != config.APIKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": fiber.StatusUnauthorized,
				"error":  "Invalid API key",
			})
		}

		return c.Next()
	}
}
