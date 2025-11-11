package logger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

// GormLogger GORM 日志适配器
type GormLogger struct {
	LogLevel logger.LogLevel
}

// NewGormLogger 创建 GORM 日志器
func NewGormLogger() *GormLogger {
	return &GormLogger{
		LogLevel: logger.Warn, // 默认只显示警告和错误
	}
}

// LogMode 设置日志级别
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 实现 logger.Interface 的 Info 方法
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		log.Info().
			Str("🗄️", "GORM").
			Interface("data", data).
			Msg(msg)
	}
}

// Warn 实现 logger.Interface 的 Warn 方法
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		log.Warn().
			Str("🗄️", "GORM").
			Interface("data", data).
			Msg(msg)
	}
}

// Error 实现 logger.Interface 的 Error 方法
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		log.Error().
			Str("🗄️", "GORM").
			Interface("data", data).
			Msg(msg)
	}
}

// Trace 实现 logger.Interface 的 Trace 方法（SQL 查询日志）
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// 只记录慢查询或错误
	if err != nil {
		log.Error().
			Str("🗄️", "GORM_SQL").
			Err(err).
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("SQL Error")
	} else if elapsed > time.Millisecond*200 { // 慢查询阈值
		log.Warn().
			Str("🗄️", "GORM_SLOW").
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("Slow SQL")
	} else if l.LogLevel <= logger.Info {
		// 只有在 Info 级别才显示所有 SQL
		log.Debug().
			Str("🗄️", "GORM_SQL").
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("SQL Query")
	}
}
