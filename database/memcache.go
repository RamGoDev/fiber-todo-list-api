package database

import (
	"fmt"
	"todo-list/configs"

	"github.com/bradfitz/gomemcache/memcache"
)

var Memcache *memcache.Client

type memcacheImpl struct {
	//
}

func NewMemcache() CacheDriver {
	return &memcacheImpl{}
}

func (impl memcacheImpl) Address() string {
	return fmt.Sprintf(
		"%s:%s",
		configs.GetEnv("MEMCACHE_HOST"),
		configs.GetEnv("MEMCACHE_PORT"),
	)
}

func (impl memcacheImpl) Connect() error {
	Memcache = memcache.New(impl.Address())

	// Test set value
	err := Memcache.Set(&memcache.Item{Key: "test_todo", Value: []byte("init value")})
	if err != nil {
		return err
	}

	return nil
}

func (impl memcacheImpl) Get(key string) string {
	item, err := Memcache.Get(key)
	if err == nil {
		return ""
	}
	return string(item.Value)
}

func (impl memcacheImpl) Set(key string, val string, sec ...int) bool {
	err := Memcache.Set(&memcache.Item{Key: key, Value: []byte(val)})
	return err == nil
}
