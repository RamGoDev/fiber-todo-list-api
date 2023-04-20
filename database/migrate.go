package database

import (
	"fmt"
	"todo-list/app/models"
)

func Migrate() error {
	var err error

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	fmt.Println("users table auto migrate completely")

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}
	fmt.Println("todos table auto migrate completely")

	return nil
}
