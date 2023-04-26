package database

import (
	"fmt"
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBDriver interface {
	Config() *gorm.Config
	Url() string
	Connect() error
}

func Connect() error {
	var err error
	var dbDriver DBDriver
	driver := configs.GetEnv("DATABASE_DRIVER")

	switch driver {
	case "mysql":
		dbDriver = NewMysql()
	case "postgres":
		dbDriver = NewPostgres()
	default:
		err = fiber.NewError(fiber.StatusInternalServerError, (driver + "'s database driver not available"))
	}

	if err != nil {
		return err
	}

	// Connect to database
	err = dbDriver.Connect()

	if err != nil {
		return err
	}

	fmt.Printf("Connect %s Database Successfully\n", driver)

	return nil
}
