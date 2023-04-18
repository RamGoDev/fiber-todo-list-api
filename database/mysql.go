package database

import (
	"fmt"
	"todo-list/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConfig() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func MysqlUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.GetEnv("DATABASE_USERNAME"),
		configs.GetEnv("DATABASE_PASSWORD"),
		configs.GetEnv("DATABASE_HOST"),
		configs.GetEnv("DATABASE_PORT"),
		configs.GetEnv("DATABASE_NAME"),
	)
}

func MysqlConnect() error {
	var err error
	url := MysqlUrl()
	DB, err = gorm.Open(mysql.Open(url), MysqlConfig())
	return err
}
