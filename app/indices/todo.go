package indices

import (
	"todo-list/app/models"
)

type Todo struct {
	ID          uint        `gorm:"primarykey" json:"id"`
	UserId      uint        `gorm:"not null" json:"user_id"`
	Title       string      `gorm:"size:150;not null" json:"title"`
	Description string      `gorm:"size:150;not null" json:"description"`
	IsDone      bool        `json:"is_done"`
	User        models.User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	models.BaseModel
}

func (u *Todo) IndexName() string {
	var todo models.Todo
	return todo.TableName()
}
