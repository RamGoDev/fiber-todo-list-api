package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Success struct {
	Error   bool        `json:"error" default:"false"`
	Data    interface{} `json:"data" default:""`
	Message string      `json:"message" default:"success"`
}

type Failed struct {
	Error   bool        `json:"error" default:"true"`
	Data    interface{} `json:"data" default:""`
	Message string      `json:"message" default:"failed"`
}

type Error struct {
	Error   bool   `json:"error" default:"true"`
	Message string `json:"message" default:"error"`
}

type Response interface {
	Success(c *fiber.Ctx, data interface{}, message string) error
	NotFound(c *fiber.Ctx, message string) error
	UnprocessableEntity(c *fiber.Ctx, data interface{}, message string) error
	TooManyRequests(c *fiber.Ctx) error
	Unauthorized(c *fiber.Ctx, message string) error
	Forbidden(c *fiber.Ctx, message string) error
}

type responseImpl struct {
	//
}

func NewResponse() Response {
	return &responseImpl{}
}

func (response responseImpl) Success(c *fiber.Ctx, data interface{}, message string) error {
	resp := Success{
		Error:   false,
		Data:    data,
		Message: message,
	}
	return c.Status(http.StatusOK).JSON(resp)
}

func (response responseImpl) NotFound(c *fiber.Ctx, message string) error {
	resp := Error{
		Error:   true,
		Message: message,
	}
	return c.Status(http.StatusNotFound).JSON(resp)
}

func (response responseImpl) UnprocessableEntity(c *fiber.Ctx, data interface{}, message string) error {
	resp := Failed{
		Error:   true,
		Data:    data,
		Message: message,
	}
	return c.Status(http.StatusUnprocessableEntity).JSON(resp)
}

func (response responseImpl) TooManyRequests(c *fiber.Ctx) error {
	resp := Failed{
		Error:   true,
		Message: "too many requests",
	}
	return c.Status(http.StatusTooManyRequests).JSON(resp)
}

func (response responseImpl) Unauthorized(c *fiber.Ctx, message string) error {
	resp := Failed{
		Error:   true,
		Message: message,
	}
	return c.Status(http.StatusUnauthorized).JSON(resp)
}

func (response responseImpl) Forbidden(c *fiber.Ctx, message string) error {
	resp := Failed{
		Error:   true,
		Message: message,
	}
	return c.Status(http.StatusForbidden).JSON(resp)
}
