package routes

import (
	V1Routes "todo-list/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MainRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New()).Name("api")

	V1Routes.RoutesV1(api)
}
