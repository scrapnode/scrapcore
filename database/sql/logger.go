package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"gorm.io/gorm/logger"
	"time"
)

type Logger struct {
	zap *xlogger.Logger
}

func (logger Logger) LogMode(logger.LogLevel) logger.Interface {
	return logger
}

func (logger Logger) Info(ctx context.Context, msg string, args ...interface{}) {
	logger.zap.Infow(msg, args...)
}
func (logger Logger) Warn(ctx context.Context, msg string, args ...interface{}) {
	logger.zap.Warnw(msg, args...)
}

func (logger Logger) Error(ctx context.Context, msg string, args ...interface{}) {
	logger.zap.Errorw(msg, args...)
}

func (logger Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()
	args := []interface{}{
		"rows", rows,
		"time", float64(elapsed.Nanoseconds()) / 1e6,
	}
	if err != nil {
		args = append(args, "error", err.Error())
	}

	logger.zap.Debugw(sql, args...)
}
