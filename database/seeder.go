package database

import (
	"errors"
	"todo-list/app/models"

	"gorm.io/gorm"
)

func UserSeeder() {
	if err := DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		DB.Create(&models.User{
			Name:     "super",
			Email:    "super@todo.id",
			Password: "password123",
		})
	}
}

func TodoSeeder() {
	if err := DB.First(&models.Todo{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var user *models.User
		err = DB.Where("id = ?", 1).Find(&user).Error

		if err != nil {
			return
		}

		var todos = []models.Todo{
			{
				Title:       "Title1",
				Description: "Description1",
				IsDone:      false,
				UserId:      user.ID,
			},
			{
				Title:       "Title2",
				Description: "Description2",
				IsDone:      true,
				UserId:      user.ID,
			},
			{
				Title:       "Title3",
				Description: "Description3",
				IsDone:      false,
				UserId:      user.ID,
			},
		}
		DB.Create(&todos)
	}
}

func DatabaseSeeder() error {
	UserSeeder()
	TodoSeeder()

	return nil
}
