package message_bus

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// NewLogger Provide zap.SugaredLogger Instance
func NewLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	config.Encoding = "json"

	zapLogger, _ := config.Build(
		zap.WithCaller(true),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return zapLogger.Sugar()
}
