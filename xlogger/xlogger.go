package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(debug bool) *Logger {
	ws := zapcore.Lock(os.Stdout)

	priority := withEnableLevel(debug)
	encoder := withEncoder(debug)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, ws, priority),
	)
	logger := zap.New(core)
	return &Logger{logger.Sugar()}
}

func withEnableLevel(debug bool) zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		if debug {
			return lvl >= zapcore.DebugLevel
		}
		return lvl >= zapcore.InfoLevel
	}
}

func withEncoder(debug bool) zapcore.Encoder {
	if debug {
		return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
