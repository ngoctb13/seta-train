package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func NewLogger(serviceName string) *Logger {
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
	infoFile, err := os.OpenFile(
		filepath.Join(serviceLogDir, "info-"+timestamp+".log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to create info log file:", err)
	}

	errorFile, err := os.OpenFile(
		filepath.Join(serviceLogDir, "error-"+timestamp+".log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to create error log file:", err)
	}

	debugFile, err := os.OpenFile(
		filepath.Join(serviceLogDir, "debug-"+timestamp+".log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to create debug log file:", err)
	}

	return &Logger{
		infoLogger:  log.New(infoFile, "[INFO] ", log.LstdFlags|log.Lshortfile),
		errorLogger: log.New(errorFile, "[ERROR] ", log.LstdFlags|log.Lshortfile),
		debugLogger: log.New(debugFile, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
	os.Exit(1)
}
