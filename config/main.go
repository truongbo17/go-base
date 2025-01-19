package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig  `mapstructure:",squash"`
	CorsConfig `mapstructure:",squash"`
}

var EnvConfig *Config

func setupConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}

func Init() {
	EnvConfig = setupConfig()
}
