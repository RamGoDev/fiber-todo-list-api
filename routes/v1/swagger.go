package routes_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "todo-list/docs"
)

func SwaggerRoutes(router fiber.Router) {
	route := router.Group("/").Name(".swagger")

	route.Get("/documentation/*", swagger.HandlerDefault).Name(".documentation")
}
