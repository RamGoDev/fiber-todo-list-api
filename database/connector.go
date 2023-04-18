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
	driver := configs.GetEnv("DATABASE_DRIVER")

	switch driver {
	case "mysql":
		err = MysqlConnect()
	default:
		err = fiber.NewError(fiber.StatusInternalServerError, (driver + "'s database driver not available"))
	}

	if err != nil {
		return err
	}

	fmt.Printf("Connect %s Database Successfully\n", driver)

	return nil
}
