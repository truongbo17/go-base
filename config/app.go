package config

type AppConfig struct {
	Env  string `mapstructure:"APP_ENV"`
	Name string `mapstructure:"APP_NAME"`
	Port string `mapstructure:"APP_PORT"`
}
