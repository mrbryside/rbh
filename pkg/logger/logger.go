package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.StacktraceKey = ""

	var err error
	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message interface{}, args ...interface{}) {
	switch v := message.(type) {
	case string:
		Log.Info(fmt.Sprintf(v, args...))
	default:
		Log.Info(fmt.Sprintf("%v", message))
	}
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Error(message interface{}, args ...interface{}) {
	switch v := message.(type) {
	case error:
		Log.Error(v.Error())

	case string:
		Log.Error(fmt.Sprintf(v, args...))
	}
}
