package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	LEVEL_TEST = -1
	LEVEL_DEV  = 0
	LEVEL_PROD = 1
)

func New(level int) *zap.SugaredLogger {
	ws := zapcore.Lock(os.Stdout)

	priority := withEnableLevel(level)
	encoder := withEncoder(level)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, ws, priority),
	)
	logger := zap.New(core)
	return logger.Sugar()
}

func withEnableLevel(level int) zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		if level == LEVEL_TEST {
			return lvl >= zapcore.WarnLevel
		}
		if level == LEVEL_DEV {
			return lvl >= zapcore.DebugLevel
		}
		return lvl >= zapcore.InfoLevel
	}
}

func withEncoder(level int) zapcore.Encoder {
	if level == LEVEL_PROD {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}
