package database

import (
	"todo-list/app/models"
)

func Migrate() error {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Todo{})

	return nil
}
