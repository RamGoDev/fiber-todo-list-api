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

	return nil
}

func (impl redisImpl) Get(key string) string {
	val, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	}
	return val
}

func (impl redisImpl) Set(key string, val string, sec ...int) bool {
	second := 0
	if len(sec) > 0 {
		second = sec[0]
	}
	err := Redis.Set(ctx, key, val, time.Duration(second)*time.Second).Err()
	return err == nil
}

func (impl redisImpl) Clear(pattern string) error {
	var foundedRecordCount int = 0

	iter := Redis.Scan(ctx, 0, pattern, 0).Iterator()
	fmt.Printf("[Redis][Clear] Your search pattern = %s\n", pattern)
	for iter.Next(ctx) {
		fmt.Printf("Deleted = %s\n", iter.Val())
		Redis.Del(ctx, iter.Val())
		foundedRecordCount++
	}

	if err := iter.Err(); err != nil {
		return err
	}

	fmt.Printf("Deleted Count %d\n", foundedRecordCount)
	return nil
}
