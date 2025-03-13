package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitializeZapLogger(debug bool) {
	var level zapcore.Level

	if debug {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		level,
	)

	Log = zap.New(core).WithOptions(zap.AddCallerSkip(1))
}

func Sync() {
	if Log != nil {
		Log.Sync()
	}
}

// Convenience functions for simpler logging
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Log.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Log.DPanic(msg, fields...)
}
