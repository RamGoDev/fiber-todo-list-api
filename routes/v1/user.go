package routes_v1

import (
	controllers_v1 "todo-list/app/controllers/v1"
	"todo-list/app/middlewares"
	repositories_v1 "todo-list/app/repositories/v1"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	route := router.Group("/users").Name(".users")
	response := helpers.NewResponse()
	repository := repositories_v1.NewUser()
	controller := controllers_v1.NewUser(response, repository)

	route.Get("/", middlewares.IsAuthenticated, middlewares.IsSuper, controller.Index).Name(".index")
	route.Post("/", middlewares.IsAuthenticated, middlewares.IsSuper, controller.Store).Name(".store")
	route.Get("/:id", middlewares.IsAuthenticated, middlewares.IsSuper, controller.Show).Name(".show")
	route.Put("/:id", middlewares.IsAuthenticated, middlewares.IsSuper, controller.Update).Name(".update")
	route.Delete("/:id", middlewares.IsAuthenticated, middlewares.IsSuper, controller.Destroy).Name(".destroy")
	route.Delete("/:id/force", middlewares.IsAuthenticated, middlewares.IsSuper, controller.ForceDestroy).Name(".destroy-force")
}
