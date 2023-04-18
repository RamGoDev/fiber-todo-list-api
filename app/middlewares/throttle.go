package middlewares

import (
	"time"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func ThrottleKeyAndIp(key string, max, sec int) func(c *fiber.Ctx) error {
	response := helpers.NewResponse()
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.OriginalURL() == "/"
		},
		Max:        max,
		Expiration: time.Duration(sec) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return string(c.IP()) + ":" + key
		},
		LimitReached: func(c *fiber.Ctx) error {
			return response.TooManyRequests(c)
		},
	})
}

func ThrottleCurrentUrlAndIp(max, sec int) func(c *fiber.Ctx) error {
	response := helpers.NewResponse()
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.OriginalURL() == "/"
		},
		Max:        max,
		Expiration: time.Duration(sec) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return string(c.IP()) + ":" + c.OriginalURL()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return response.TooManyRequests(c)
		},
	})
}
