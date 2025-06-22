package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"jubeliotesting/pkg/logger"
	"time"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next() // eksekusi handler berikutnya

		stop := time.Now()
		latency := stop.Sub(start)

		logger.Log.Info("API Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
			zap.String("user-agent", c.Get("User-Agent")),
		)

		return err
	}
}
