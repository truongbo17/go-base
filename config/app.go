package config

type AppConfig struct {
	Name string `mapstructure:"APP_NAME" default:"GoGinBase"`
	Port string `mapstructure:"APP_PORT" default:"8000"`
}
