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
	Set(key string, val string, sec ...int) bool

	// Clear cache with pattern
	Clear(pattern string) error
}

func GetCacheDriver() (CacheDriver, string, error) {
	var err error
	var cacheDriver CacheDriver
	driver := configs.GetEnv("CACHE_DRIVER")

	switch driver {
	case "redis":
		cacheDriver = NewRedis()
	case "memcache":
		cacheDriver = NewMemcache()
	default:
		err = fiber.NewError(fiber.StatusInternalServerError, (driver + "'s cache driver not available"))
	}

	if err != nil {
		return nil, "", err
	}

	return cacheDriver, driver, nil
}

func CacheConnect() error {
	cacheDriver, driver, err := GetCacheDriver()

	if err != nil {
		return err
	}

	// Connect to cache client
	err = cacheDriver.Connect()

	if err != nil {
		return err
	}

	fmt.Printf("Connect %s's Cache Driver Successfully\n", driver)

	return nil
}
