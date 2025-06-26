package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func OpenTelemetryMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tracer := otel.Tracer("myapp")
		_, span := tracer.Start(c.Context(), c.Method()+" "+c.Path())
		c.Locals("otel-span", span)
		defer span.End()
		err := c.Next()
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
		} else {
			span.SetStatus(codes.Ok, "OK")
		}
		return err
	}
}
