package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	Logger *zerolog.Logger
}

func InitLogger(serviceName string) *Logger {
	// Create logs directory if it doesn't exist
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}

	// Create service-specific log directory
	serviceLogDir := filepath.Join(logsDir, serviceName)
	if err := os.MkdirAll(serviceLogDir, 0755); err != nil {
		log.Fatal("Failed to create service log directory:", err)
	}

	// Create log files with timestamp
	timestamp := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(
		filepath.Join(serviceLogDir, "log-"+timestamp+".log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		log.Fatal("Failed to create info log file:", err)
	}

	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.Logger.Info().Msgf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.Logger.Error().Msgf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.Logger.Debug().Msgf(format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Logger.Fatal().Msgf(format, v...)
	os.Exit(1)
}
