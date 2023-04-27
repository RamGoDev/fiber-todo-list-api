package indices

import (
	"todo-list/app/models"
)

type User struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `gorm:"type:varchar(150);not null" json:"name"`
	Email string `gorm:"type:varchar(100);unique;not null" json:"email"`
	models.BaseModel
}

func (u *User) IndexName() string {
	var user models.User
	return user.TableName()
}
