package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
}

func Debug(msg string, tags ...zap.Field) {
	log.Debug(msg, tags...)
}

func Error(msg string, tags ...zap.Field) {
	log.Error(msg, tags...)
}

func Fatal(msg string, tags ...zap.Field) {
	log.Fatal(msg, tags...)
}
