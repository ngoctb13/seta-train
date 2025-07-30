package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/shared-modules/utils"
)

type LoggingMiddleware struct {
	logger *utils.Logger
}

func NewLoggingMiddleware(serviceName string) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: utils.NewLogger(serviceName),
	}
}

func (lm *LoggingMiddleware) LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log details
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		bodySize := c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		lm.logger.Info("[%s] %s %s %d %d %v %s",
			method,
			path,
			clientIP,
			statusCode,
			bodySize,
			latency,
			c.Errors.String(),
		)

		// Log errors separately
		if len(c.Errors) > 0 {
			lm.logger.Error("Request errors: %s", c.Errors.String())
		}
	}
}

func (lm *LoggingMiddleware) ErrorLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Log any errors that occurred during request processing
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				lm.logger.Error("Request error: %s - %s", c.Request.URL.Path, err.Error())
			}
		}
	}
}
