package logger

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger *zap.Logger) Logger {
	return Logger{logger: logger}
}

func (l Logger) Middleware(c *gin.Context) {
	var requestBody string
	if c.Request != nil {
		body, marshalErr := io.ReadAll(c.Request.Body)
		if marshalErr != nil {
			l.logger.Error(marshalErr.Error(),
				zap.String("type", "Logger Middleware error"),
			)
		}

		requestBody = string(body)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	}
	c.Next()

	// err := models.APILog{
	// 	Datetime: time.Now(),
	// 	IPFrom:   c.ClientIP(),
	// 	Method:   c.Request.Method,
	// 	Status:   c.Writer.Status(),
	// 	URL:      c.Request.RequestURI,
	// 	Request:  string(body),
	// }.Create()
	// if err != nil {
	// 	logger.ZapLogger.Error(err.Error(),
	// 		zap.String("type", "Mysql service error"),
	// 	)
	// }

	l.logger.Info("Middleware logger",
		zap.String("IP_From", c.ClientIP()),
		zap.String("method", c.Request.Method),
		zap.Int("status", c.Writer.Status()),
		zap.String("url", c.Request.RequestURI),
		zap.String("request", requestBody),
	)
	c.Next()
}
