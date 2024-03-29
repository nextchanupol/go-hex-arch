package hlps

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var err error

func init() {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	// log.Error(message, fields...)
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v)
	}
}
