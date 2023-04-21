package database

import (
	"fmt"
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
)

type CacheDriver interface {
	// Connect to cache client
	Connect() error

	// Get cache value with key
	Get(key string) string

	// Set cache with second duration
	Set(key string, val string, sec int) bool
}

func CacheConnect() error {
	var err error
	var cacheDriver CacheDriver
	driver := configs.GetEnv("CACHE_DRIVER")

	switch driver {
	case "redis":
		cacheDriver = NewRedis()
	default:
		err = fiber.NewError(fiber.StatusInternalServerError, (driver + "'s cache driver not available"))
	}

	if err != nil {
		return err
	}

	// Connect to cache client
	cacheDriver.Connect()

	fmt.Printf("Connect %s's Cache Driver Successfully\n", driver)

	return nil
}
