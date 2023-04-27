package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"todo-list/app/indices"
	"todo-list/app/models"
	"todo-list/helpers"

	"gorm.io/gorm"
)

func UserSeeder() {
	if err := DB.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var user = models.User{
			Name:     "super",
			Email:    "super@todo.id",
			Password: "password123",
		}
		var userIndex *indices.User

		DB.Create(&user)

		// Store to elasticsearch
		elastic := NewElasticsearch()

		// Delete index for reset
		elastic.DeleteIndex([]string{userIndex.IndexName()})

		_, _ = helpers.ConvertToOtherStruct(user, &userIndex)
		dataByte, _ := json.Marshal(userIndex)
		elastic.AddDocument(userIndex.IndexName(), dataByte)

		fmt.Println("users table sedder successfully")
	}
}

func TodoSeeder() {
	if err := DB.First(&models.Todo{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		var user *models.User
		var todoIndex *indices.Todo

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

		// Store to elasticsearch
		elastic := NewElasticsearch()

		// Delete index for reset
		elastic.DeleteIndex([]string{todoIndex.IndexName()})

		var dataByte []byte
		for _, todo := range todos {
			_, _ = helpers.ConvertToOtherStruct(todo, &todoIndex)
			dataByte, _ = json.Marshal(todoIndex)
			elastic.AddDocument(todoIndex.IndexName(), dataByte)
		}

		fmt.Println("todos table sedder successfully")
	}
}

func DatabaseSeeder() error {
	UserSeeder()
	TodoSeeder()

	return nil
}
