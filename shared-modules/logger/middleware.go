package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

type LoggingMiddleware struct {
	Logger *Logger
}

func NewLoggingMiddleware(logger *Logger, serviceName string) *LoggingMiddleware {
	return &LoggingMiddleware{
		Logger: logger,
	}
}

func (lm *LoggingMiddleware) LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		lm.Logger.Logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Msg("Incoming request")
	}
}
