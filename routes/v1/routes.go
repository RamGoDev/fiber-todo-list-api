package routes_v1

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesV1(api fiber.Router) {
	V1Routes := api.Group("/v1").Name(".v1")

	TestRoutes(V1Routes)
	UserRoutes(V1Routes)
	AuthRoutes(V1Routes)
	TodoRoutes(V1Routes)
	SwaggerRoutes(V1Routes)
}
