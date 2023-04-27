package models

import (
	"fmt"
)

type Todo struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	UserId      uint   `gorm:"not null" json:"user_id"`
	Title       string `gorm:"size:150;not null" json:"title"`
	Description string `gorm:"size:150;not null" json:"description"`
	IsDone      bool   `json:"is_done"`
	User        User   `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BaseModel
}

func (u *Todo) TableName() string {
	return "todos"
}

func (u *Todo) CacheBaseKey() string {
	return u.TableName()
}

func (u *Todo) CacheShowKey(userId string, id string) string {
	return fmt.Sprintf("%s_%s_%s", userId, u.CacheBaseKey(), id)
}
