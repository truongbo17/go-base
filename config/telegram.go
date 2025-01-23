package config

const (
	URLBaseTelegram     string = "https://api.telegram.org/bot%s"
	EndpointSendMessage string = "/sendMessage"
)

type TelegramConfig struct {
	BotToken      string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	ChatIdError   string `mapstructure:"TELEGRAM_CHAT_ID_ERROR"`
	ChatIdWarning string `mapstructure:"TELEGRAM_CHAT_ID_WARNING"`
}
