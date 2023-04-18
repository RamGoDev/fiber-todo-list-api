package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	UserId      uint   `gorm:"not null" json:"user_id"`
	Title       string `gorm:"size:150;not null" json:"title"`
	Description string `gorm:"size:150;not null" json:"description"`
	IsDone      bool   `json:"is_done"`
	User        User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}

func (u *Todo) TableName() string {
	return "todos"
}
