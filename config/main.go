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

	viper.AddConfigPath(path + "/config")
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

	fmt.Println(config.AppConfig.Port)

	return config
}

func Init() {
	EnvConfig = setupConfig()
}
