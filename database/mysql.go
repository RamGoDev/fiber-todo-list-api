package database

import (
	"fmt"
	"todo-list/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlImpl struct {
	//
}

func NewMysql() DBDriver {
	return &mysqlImpl{}
}

func (impl mysqlImpl) Config() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	}
}

func (impl mysqlImpl) Url() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.GetEnv("DATABASE_USERNAME"),
		configs.GetEnv("DATABASE_PASSWORD"),
		configs.GetEnv("DATABASE_HOST"),
		configs.GetEnv("DATABASE_PORT"),
		configs.GetEnv("DATABASE_NAME"),
	)
}

func (impl mysqlImpl) Connect() error {
	var err error
	url := impl.Url()
	DB, err = gorm.Open(mysql.Open(url), impl.Config())
	return err
}
