package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func JsonHeader(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Next()
}
