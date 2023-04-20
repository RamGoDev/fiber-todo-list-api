package database

import (
	"fmt"
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dbDriver := configs.GetEnv("DATABASE_DRIVER")

	switch dbDriver {
	case "mysql":
		err = NewMysql().Connect()
	case "postgres":
		err = NewPostgres().Connect()
	default:
		err = fiber.NewError(fiber.StatusInternalServerError, (dbDriver + "'s database driver not available"))
	}

	if err != nil {
		return err
	}

	fmt.Printf("Connect %s Database Successfully\n", dbDriver)

	return nil
}
