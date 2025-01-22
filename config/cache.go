package config

const (
	CacheStoreLocal   string = "local"
	CacheStoreRedis   string = "redis"
	CacheKeyRateLimit string = "router_rate_limit"
)

type CacheConfig struct {
	CacheStore    string `mapstructure:"CACHE_STORE"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisUsername string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}
