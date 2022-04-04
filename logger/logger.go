package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	// zapConfig := zap.NewProductionConfig()
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logger, err = zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.Warn(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch myType := message.(type) {
	case error:
		logger.Error(myType.Error(), fields...)
	case string:
		logger.Error(myType)
	}
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}
