package middlewares

import (
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	response := helpers.NewResponse()
	valid, err := helpers.ValidateJWT(c)

	if !valid {
		return response.Unauthorized(c, err.Error())
	}

	return c.Next()
}

func IsSuper(c *fiber.Ctx) error {
	response := helpers.NewResponse()
	userId := helpers.GetCurrentUserId(c)

	// Note: use 1 for testing
	if userId != "1" {
		return response.Unauthorized(c, "you don't have permission")
	}

	return c.Next()
}
