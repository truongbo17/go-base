package config

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig          `mapstructure:",squash"`
	CorsConfig         `mapstructure:",squash"`
	DatabaseConnection `mapstructure:",squash"`
	CacheConfig        `mapstructure:",squash"`
	AuthConfig         `mapstructure:",squash"`
}

const (
	// DebugMode app env debug stg.
	DebugMode string = "debug"
	// ReleaseMode app env debug production.
	ReleaseMode string = "release"
)

func (config *Config) validate() error {
	return validation.ValidateStruct(config,
		// App
		validation.Field(&config.AppConfig.Port, is.Port),
		validation.Field(&config.AppConfig.Env, validation.In(DebugMode, ReleaseMode)),
		validation.Field(&config.AppConfig.IsWorker, validation.In(true, false)),

		// CORS
		validation.Field(&config.CorsConfig.AllowOrigin),

		// Database
		validation.Field(&config.DatabaseConnection.DatabaseRelation.Port, is.Port),
		validation.Field(&config.DatabaseConnection.DatabaseRelation.Host, is.Host),

		// Cache
		validation.Field(&config.CacheConfig.CacheStore, validation.In(CacheStoreLocal, CacheStoreRedis)),

		// Redis
		validation.Field(&config.CacheConfig.RedisPort, is.Port),
		validation.Field(&config.CacheConfig.RedisHost, is.Host),

		// Auth
		validation.Field(&config.AuthConfig.JWTSecretKey, validation.Required),
	)
}

var EnvConfig *Config

func setupConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetDefault("APP_ENV", "debug")
	viper.SetDefault("APP_PORT", "8000")
	viper.SetDefault("CORS_ALLOW_ORIGIN", "*")
	viper.SetDefault("CACHE_STORE", "local")

	viper.SetDefault("JWT_ACCESS_EXPIRATION_MINUTES", 120)
	viper.SetDefault("JWT_REFRESH_EXPIRATION_DAYS", 7)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	if err := config.validate(); err != nil {
		panic(err)
	}

	return config
}

func Init() {
	EnvConfig = setupConfig()
}
