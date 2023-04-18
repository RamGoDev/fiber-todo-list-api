package middlewares

import (
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func DefaultMiddlewares(app *fiber.App) {
	app.Use(
		logger.New(),
		JsonHeader,
		ThrottleCurrentUrlAndIp(configs.GetEnvInt("RATE_LIMITER_MAX"), configs.GetEnvInt("RATE_LIMITER_TIME")),
	)
}
