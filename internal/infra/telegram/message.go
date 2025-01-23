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

type ResponseTelegram struct {
	Ok     bool           `json:"ok"`
	Result ResultTelegram `json:"result,omitempty"`
}

type ResultTelegram struct {
	MessageID int          `json:"message_id"`
	From      UserTelegram `json:"from"`
	Chat      ChatTelegram `json:"chat"`
	Date      int64        `json:"date"`
	Text      string       `json:"text"`
}

type UserTelegram struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type ChatTelegram struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

func SendMessage(chatID string, text string) {
	logApp := logger.LogrusLogger

	botToken := config.EnvConfig.TelegramConfig.BotToken
	url := fmt.Sprintf(config.URLBaseTelegram+config.EndpointSendMessage, botToken)
	message := MessageStruct{
		ChatID: chatID,
		Text:   text,
	}

	client := httpClient.NewBaseRequest(10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	_, body, err := client.Post(ctx, url, headers, message)
	if err != nil {
		logApp.Errorln("Send message to telegram error: ", err)
		return
	}

	var telegramResponse ResponseTelegram
	_ = json.Unmarshal(body, &telegramResponse)
	logApp.Infof("Telegram send message response: %+v", string(body))
}
