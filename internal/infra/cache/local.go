package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type LocalCache struct {
	store *cache.Cache
}

func NewLocalCache(defaultExpiration, cleanupInterval time.Duration) *LocalCache {
	return &LocalCache{
		store: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (l *LocalCache) Set(key string, value interface{}, ttl time.Duration) error {
	l.store.Set(key, value, ttl)
	return nil
}

func (l *LocalCache) Get(key string) (interface{}, error) {
	val, found := l.store.Get(key)
	if !found {
		return nil, nil
	}
	return val, nil
}

func (l *LocalCache) Delete(key string) error {
	l.store.Delete(key)
	return nil
}
