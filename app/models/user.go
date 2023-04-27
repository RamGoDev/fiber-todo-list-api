package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"type:varchar(150);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"size:150;not null" json:"password"`
	gorm.Model
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) CacheBaseKey() string {
	return u.TableName()
}

func (u *User) CacheShowKey(id string) string {
	return fmt.Sprintf("%s_%s", u.CacheBaseKey(), id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	HashPassword, err := hashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = HashPassword

	return nil
}
