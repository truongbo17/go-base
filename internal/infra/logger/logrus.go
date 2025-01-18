package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

func InitLogrusLogger() *logrus.Logger {
	logger := logrus.New()
	currentDate := time.Now()

	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logFile := &lumberjack.Logger{
		Filename: fmt.Sprintf("logs/%s.log", currentDate.Format("2006-01-02")),
		MaxSize:  10,
		Compress: false,
	}

	writers := []io.Writer{logFile, os.Stdout}
	logger.SetOutput(io.MultiWriter(writers...))

	return logger
}
