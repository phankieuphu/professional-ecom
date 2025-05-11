package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		entry := logger.WithFields(logrus.Fields{
			"status":  statusCode,
			"method":  method,
			"path":    path,
			"query":   raw,
			"ip":      clientIP,
			"latency": latency,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("Reuqest handled")
		}
	}
}
