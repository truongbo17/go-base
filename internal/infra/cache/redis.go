package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(context.Background(), key, value, ttl).Err()
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	return val, err
}

func (r *RedisCache) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
