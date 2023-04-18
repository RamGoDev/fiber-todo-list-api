package main

import (
	"fmt"
	"todo-list/app/middlewares"
	"todo-list/database"
	"todo-list/helpers"
	"todo-list/routes"

	"github.com/gofiber/fiber/v2"
)

func homepage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	apiDocumentUrl, _ := c.GetRouteURL("api.v1.swagger.documentation", fiber.Map{})
	html := `<h1>Welcome to Todo List API !</h1>`
	html = html + fmt.Sprintf("Go to <a href='%s'>%s</a> for API Documentation", apiDocumentUrl, apiDocumentUrl)
	return c.SendString(html)
}

// @title Todo List API
// @version 1.0
// @description This is a Todo List API swagger
// @contact.name Ramdani
// @contact.url https://github.com/ramdani15/
// @contact.email ramdaninformatika@gmail.com
// @BasePath /
// @schemas http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := fiber.New()

	// Connect Database
	errDb := database.Connect()
	if errDb != nil {
		panic("Failed to connect Database (" + errDb.Error() + ")")
	}

	// Connect Redis for Caching
	errCache := database.RedisConnect()
	if errCache != nil {
		panic("Failed to connect Redis (" + errCache.Error() + ")")
	}

	// Migrate
	errMigrate := database.Migrate()
	if errMigrate != nil {
		panic(errMigrate.Error())
	}

	// Seeder
	errSeeder := database.DatabaseSeeder()
	if errSeeder != nil {
		panic(errSeeder.Error())
	}

	// Set Middleware
	middlewares.DefaultMiddlewares(app)

	// Routes
	routes.MainRoutes(app)

	app.Get("/", homepage).Name("homepage")

	// Run Server
	helpers.StartServer(app)
}
