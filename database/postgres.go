package database

import (
	"fmt"
	"todo-list/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresImpl struct {
	//
}

func NewPostgres() DBDriver {
	return &postgresImpl{}
}

func (impl postgresImpl) Config() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	}
}

func (impl postgresImpl) Url() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.GetEnv("DATABASE_HOST"),
		configs.GetEnv("DATABASE_USERNAME"),
		configs.GetEnv("DATABASE_PASSWORD"),
		configs.GetEnv("DATABASE_NAME"),
		configs.GetEnv("DATABASE_PORT"),
	)
}

func (impl postgresImpl) Connect() error {
	var err error
	url := impl.Url()
	DB, err = gorm.Open(postgres.Open(url), impl.Config())
	return err
}
