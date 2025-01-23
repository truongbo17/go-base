package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"go-base/config"
	httpClient "go-base/internal/infra/http"
	"go-base/internal/infra/logger"
	"time"
)

type MessageStruct struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(chatID string, text string) {
	logApp := logger.LogrusLogger

	botToken := config.EnvConfig.TelegramConfig.BotToken
	url := fmt.Sprintf(config.URLBaseTelegram+config.EndpointSendMessage, botToken)
	message := MessageStruct{
		ChatID: chatID,
		Text:   text,
	}

	messageBody, err := json.Marshal(message)
	if err != nil {
		logApp.Errorf("failed to marshal message: %w", err)
		return
	}

	client := httpClient.NewBaseRequest(10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	headers := map[string]string{
		"Accept": "application/json",
	}
	_, _, err = client.Post(ctx, url, headers, messageBody)
	if err != nil {
		logApp.Errorf("Send message to telegram error: %w", err)
		return
	}
}
