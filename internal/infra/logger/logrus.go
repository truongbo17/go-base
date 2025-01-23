package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-base/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

func InitLogrusLogger() *logrus.Logger {
	logger := logrus.New()
	currentDate := time.Now()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.DateTime,
		DisableQuote:    true,
	})

	logFile := &lumberjack.Logger{
		Filename: fmt.Sprintf(config.PathLog, currentDate.Format(time.DateOnly)),
		MaxSize:  10,
		Compress: false,
	}

	writers := []io.Writer{logFile, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	logger.Infoln("Success init logger with Logrus")

	return logger
}
