package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	"go-base/config"
	"go-base/internal/infra/logger"
	"time"
)

type StoreCache interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string, value interface{}) (interface{}, error)
	Delete(key string) error
}

type RedisCache struct {
	store *persist.RedisStore
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	return r.store.Set(key, value, ttl)
}

func (r *RedisCache) Get(key string, value interface{}) (interface{}, error) {
	return r.store.Get(key, value), nil
}

func (r *RedisCache) Delete(key string) error {
	return r.store.Delete(key)
}

type MemoryCache struct {
	store *persist.MemoryStore
}

func (m *MemoryCache) Set(key string, value interface{}, ttl time.Duration) error {
	return m.store.Set(key, value, ttl)
}

func (m *MemoryCache) Get(key string, value interface{}) (interface{}, error) {
	return m.store.Get(key, value), nil
}

func (m *MemoryCache) Delete(key string) error {
	return m.store.Delete(key)
}

var Cache StoreCache

func InitCache(store string) {
	if store == config.CacheStoreRedis {
		Cache = &RedisCache{store: InitCacheRedis()}
	} else {
		Cache = &MemoryCache{store: InitCacheLocal()}
	}
	logApp := logger.LogrusLogger
	logApp.Infoln("Success init cache with store " + store)
}
