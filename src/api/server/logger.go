package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
)

var replacer = strings.NewReplacer("\t", " ", "\n", " ")

type Logger struct {
	*zap.SugaredLogger
}

type GinLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

// ProvideLogger Provide Logger Instance
func ProvideLogger() *Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	config.Encoding = "json"

	zapLogger, _ := config.Build(
		zap.WithCaller(true),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	logger := NewSugaredLogger(zapLogger)
	return logger
}

// NewSugaredLogger Generate a sugared logger
func NewSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetGinLogger Get gin logger
func (l Logger) GetGinLogger() *GinLogger {
	return &GinLogger{
		Logger: &l,
	}
}

// Write Override gin-framework method logger
func (l GinLogger) Write(p []byte) (n int, err error) {
	message := replacer.Replace(string(p))
	l.Info(message)
	return len(p), nil
}

// GetFxLogger Get fx logger
func (l Logger) GetFxLogger() *FxLogger {
	return &FxLogger{
		Logger: &l,
	}
}

// Printf Override fx method logger
func (l FxLogger) Printf(str string, args ...interface{}) {
	message := replacer.Replace(str)
	if args != nil {
		l.Infof(message, args)
	} else {
		l.Info(message)
	}
}
