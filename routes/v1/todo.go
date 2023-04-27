package routes_v1

import (
	controllers_v1 "todo-list/app/controllers/v1"
	"todo-list/app/middlewares"
	repositories_v1 "todo-list/app/repositories/v1"
	validators_v1 "todo-list/app/validators/v1"
	"todo-list/database"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(router fiber.Router) {
	route := router.Group("/todos").Name(".todos")
	response := helpers.NewResponse()
	cache, _, _ := database.GetCacheDriver()
	elastic := database.NewElasticsearch()
	repository := repositories_v1.NewTodo(cache, elastic)
	controller := controllers_v1.NewTodo(response, repository, cache)

	route.Get("/", middlewares.IsAuthenticated, controller.Index).Name(".index")
	route.Post("/", middlewares.IsAuthenticated, validators_v1.TodoValidator, controller.Store).Name(".store")
	route.Get("/:id", middlewares.IsAuthenticated, controller.Show).Name(".show")
	route.Put("/:id", middlewares.IsAuthenticated, validators_v1.TodoValidator, controller.Update).Name(".update")
	route.Delete("/:id", middlewares.IsAuthenticated, controller.Destroy).Name(".destroy")
	route.Delete("/:id/force", middlewares.IsAuthenticated, controller.ForceDestroy).Name(".destroy-force")
	route.Put("/:id/complete", middlewares.IsAuthenticated, controller.Complete).Name(".complete")
}
