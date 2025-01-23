package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-base/internal/infra/logger"
)

const TypeEmailRegister = "email_register"

type SendEmailRegisterPayload struct {
	UserID int
	Email  string
}

func SendMailRegisterTask(userId uint, email string) (*asynq.Task, error) {
	logApp := logger.LogrusLogger
	logApp.Infof("Start send email register task with userId: %d, and email: %s", userId, email)
	payload, err := json.Marshal(SendEmailRegisterPayload{UserID: int(userId), Email: email})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailRegister, payload), nil
}

func HandleSendMailRegister(ctx context.Context, t *asynq.Task) error {
	var payload SendEmailRegisterPayload
	logApp := logger.LogrusLogger

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	logApp.Infof("Success send email register task with userId: %d, and email: %s", payload.UserID, payload.Email)
	return nil
}
