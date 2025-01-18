package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	AppConfig AppConfig
}

var EnvConfig *Config

func setupConfig() *Config {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigFile(".env")
	viper.AddConfigPath(path + "/config")

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
