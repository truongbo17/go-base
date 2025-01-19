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
}

const (
	// DebugMode app env debug stg.
	DebugMode = "debug"
	// ReleaseMode app env debug production.
	ReleaseMode = "release"
)

func (config *Config) validate() error {
	return validation.ValidateStruct(config,
		validation.Field(&config.AppConfig.Port, is.Port),
		validation.Field(&config.AppConfig.Env, validation.In(DebugMode, ReleaseMode)),

		//validation.Field(&config.UseRedis, validation.In(true, false)),
		//validation.Field(&config.RedisDefaultAddr),
		//
		//validation.Field(&config.JWTSecretKey, validation.Required),
		//validation.Field(&config.JWTAccessExpirationMinutes, validation.Required),
		//validation.Field(&config.JWTRefreshExpirationDays, validation.Required),
	)
}

var EnvConfig *Config

func setupConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetDefault("APP_ENV", "debug")
	viper.SetDefault("APP_PORT", "8000")
	viper.SetDefault("CORS_ALLOW_ORIGIN", "*")

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
