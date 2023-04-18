package helpers

import (
	"fmt"
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
)

func getServerHost() string {
	return fmt.Sprintf(
		"%s:%s",
		configs.GetEnv("APP_HOST"),
		configs.GetEnv("APP_PORT"),
	)
}

func StartServer(app *fiber.App) {
	host := getServerHost()
	server := app.Listen(host)

	if server != nil {
		panic(server.Error())
	}
}
