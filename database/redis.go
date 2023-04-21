package database

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"todo-list/configs"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var Redis *redis.Client

type redisImpl struct {
	//
}

func NewRedis() CacheDriver {
	return &redisImpl{}
}

func (impl redisImpl) Address() string {
	return fmt.Sprintf(
		"%s:%s",
		configs.GetEnv("REDIS_HOST"),
		configs.GetEnv("REDIS_PORT"),
	)
}

func (impl redisImpl) Options() *redis.Options {
	redisDb, _ := strconv.Atoi(configs.GetEnv("REDIS_DB"))
	return &redis.Options{
		Addr:     impl.Address(),
		Password: configs.GetEnv("REDIS_PASSWORD"),
		DB:       redisDb,
	}
}

func (impl redisImpl) Connect() error {
	Redis = redis.NewClient(impl.Options())

	// Test set value
	err := Redis.Set(ctx, "test_todo", "init value", 0).Err()
	if err != nil {
		return err
	}

	fmt.Println("Connect Redis Successfully")

	return nil
}

func (impl redisImpl) Get(key string) string {
	val, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	}
	return val
}

func (impl redisImpl) Set(key string, val string, sec int) bool {
	err := Redis.Set(ctx, key, val, time.Duration(sec)*time.Second).Err()
	return err == nil
}
