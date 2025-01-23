package config

const (
	HeaderAuth string = "Authorization"
	TokenType  string = "Bearer"
)

type AuthConfig struct {
	JWTSecretKey               string `mapstructure:"JWT_SECRET"`
	JWTAccessExpirationMinutes int    `mapstructure:"JWT_ACCESS_EXPIRATION_MINUTES"`
	JWTRefreshExpirationDays   int    `mapstructure:"JWT_REFRESH_EXPIRATION_DAYS"`
	GoogleAuthClientID         string `mapstructure:"AUTH_SOCIAL_GOOGLE_CLIENT_ID"`
	GoogleAuthClientSecret     string `mapstructure:"AUTH_SOCIAL_GOOGLE_CLIENT_SECRET"`
}
