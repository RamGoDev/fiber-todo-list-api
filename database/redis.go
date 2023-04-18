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

func GetRedisAddress() string {
	return fmt.Sprintf(
		"%s:%s",
		configs.GetEnv("REDIS_HOST"),
		configs.GetEnv("REDIS_PORT"),
	)
}

func RedisOptions() *redis.Options {
	redisDb, _ := strconv.Atoi(configs.GetEnv("REDIS_DB"))
	return &redis.Options{
		Addr:     GetRedisAddress(),
		Password: configs.GetEnv("REDIS_PASSWORD"),
		DB:       redisDb,
	}
}

// Connect to redis client
func RedisConnect() error {
	Redis = redis.NewClient(RedisOptions())

	// Test set value
	err := Redis.Set(ctx, "test_todo", "init value", 0).Err()
	if err != nil {
		return err
	}

	fmt.Println("Connect Redis Successfully")

	return nil
}

// Get cache value with key
//
// param	key	string
// return	string
func RedisGet(key string) string {
	val, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	}
	return val
}

// Set cache with second duration
//
// param	key	string
// param	value	string
// param	sec	int
// return	string
func RedisSet(key string, val string, sec int) bool {
	err := Redis.Set(ctx, key, val, time.Duration(sec)*time.Second).Err()
	return err == nil
}
