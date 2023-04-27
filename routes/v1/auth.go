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

func AuthRoutes(router fiber.Router) {
	route := router.Group("/auth").Name(".auth")
	response := helpers.NewResponse()
	cache, _, _ := database.GetCacheDriver()
	elastic := database.NewElasticsearch()
	repository := repositories_v1.NewAuth(elastic)
	userRepository := repositories_v1.NewUser(cache, elastic)
	controller := controllers_v1.NewAuth(response, repository, userRepository)

	route.Post("/login", validators_v1.LoginValidator, controller.Login).Name(".login")
	route.Post("/register", validators_v1.RegisterValidator, controller.Register).Name(".register")
	route.Get("/profile", middlewares.IsAuthenticated, controller.MyProfile).Name(".my-profile")
	route.Put("/profile", middlewares.IsAuthenticated, validators_v1.ProfileValidator, controller.UpdateProfile).Name(".update-profile")
}
