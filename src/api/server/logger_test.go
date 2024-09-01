package server_test

import (
	"api.ddd/src/api/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestWhenCallProvideLoggerShouldReturnLogger(t *testing.T) {
	// Act
	logger := server.ProvideLogger()

	// Assert
	assert.NotNil(t, logger)
	assert.IsType(t, &server.Logger{}, logger)
	assert.NotNil(t, logger.SugaredLogger)
}

func TestGivenZapLoggerWhenCallNewSugaredLoggerShouldReturnLogger(t *testing.T) {
	// Arrange
	zapLogger, _ := zap.NewDevelopmentConfig().Build()

	// Act
	logger := server.NewSugaredLogger(zapLogger)

	// Assert
	assert.NotNil(t, logger)
	assert.IsType(t, &server.Logger{}, logger)
	assert.NotNil(t, logger.SugaredLogger)
}

func TestGivenLoggerWhenCallGetGinLoggerShouldReturnGinLogger(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	ginLogger := logger.GetGinLogger()

	// Assert
	assert.NotNil(t, ginLogger)
	assert.IsType(t, &server.GinLogger{}, ginLogger)
	assert.NotNil(t, ginLogger.SugaredLogger)
}

func TestGivenGinLoggerWhenCallWriteShouldReturnNilError(t *testing.T) {
	// Arrange
	ginLogger := server.ProvideLogger().GetGinLogger()
	message := "hello world"
	messageLen := len(message)
	messageBytes := []byte(message)

	// Act
	length, err := ginLogger.Write(messageBytes)

	// Assert
	assert.NotNil(t, length)
	assert.Equal(t, messageLen, length)
	assert.Nil(t, err)
}

func TestGivenFxLoggerWhenCallGetFxLoggerShouldReturnFxLogger(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger()

	// Act
	fxLogger := logger.GetFxLogger()

	// Assert
	assert.NotNil(t, fxLogger)
	assert.IsType(t, &server.FxLogger{}, fxLogger)
	assert.NotNil(t, fxLogger.SugaredLogger)
}

func TestGivenFxLoggerWhenCallPrintfWithoutArgsShouldCallInfo(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger().GetFxLogger()

	// Act
	logger.Printf("dummy-message")

	// Assert
	assert.NotNil(t, logger)
	assert.IsType(t, &server.FxLogger{}, logger)
	assert.NotNil(t, logger.SugaredLogger)
}

func TestGivenFxLoggerWhenCallPrintfWithArgsShouldCallInfof(t *testing.T) {
	// Arrange
	logger := server.ProvideLogger().GetFxLogger()

	// Act
	logger.Printf("dummy-message", []interface{}{"arg"})

	// Assert
	assert.NotNil(t, logger)
	assert.IsType(t, &server.FxLogger{}, logger)
	assert.NotNil(t, logger.SugaredLogger)
}
