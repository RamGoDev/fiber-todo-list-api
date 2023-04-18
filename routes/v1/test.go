package routes_v1

import (
	V1Controllers "todo-list/app/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(router fiber.Router) {
	route := router.Group("/test")

	route.Get("/", V1Controllers.Test)
	route.Post("/", V1Controllers.TestPost)
}
