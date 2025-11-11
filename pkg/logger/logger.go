package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var log zerolog.Logger

// 初始化日志系统
func init() {
	// 创建彩色控制台输出
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		// 自定义格式
		FormatLevel: func(i interface{}) string {
			var level string
			if ll, ok := i.(string); ok {
				switch ll {
				case "debug":
					level = "🐛 DEBUG"
				case "info":
					level = "ℹ️  INFO"
				case "warn":
					level = "⚠️  WARN"
				case "error":
					level = "❌ ERROR"
				case "fatal":
					level = "💀 FATAL"
				case "panic":
					level = "🚨 PANIC"
				default:
					level = "📝 " + ll
				}
			} else {
				level = "📝 " + "???"
			}
			return level
		},
		FormatMessage: func(i interface{}) string {
			return "| " + i.(string)
		},
	}

	log = zerolog.New(consoleWriter).
		With().
		Timestamp().
		Logger()
}

// ==================== 基础日志方法 ====================

func Debug() *zerolog.Event {
	return log.Debug()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Warn() *zerolog.Event {
	return log.Warn()
}

func Error() *zerolog.Event {
	return log.Error()
}

func Fatal() *zerolog.Event {
	return log.Fatal()
}

// ==================== 快捷方法 ====================

// Success 成功日志
func Success(message string) {
	log.Info().Str("💚", "SUCCESS").Msg(message)
}

// Fail 失败日志
func Fail(message string) {
	log.Error().Str("💔", "FAIL").Msg(message)
}

// Start 开始操作
func Start(operation string) {
	log.Info().Str("🚀", "START").Msg(operation)
}

// Done 完成操作
func Done(operation string) {
	log.Info().Str("✅", "DONE").Msg(operation)
}

// Connecting 连接中
func Connecting(service string) {
	log.Info().Str("🔌", "CONNECTING").Msg(service)
}

// Connected 已连接
func Connected(service string) {
	log.Info().Str("🔗", "CONNECTED").Msg(service)
}

// Loading 加载中
func Loading(resource string) {
	log.Info().Str("📥", "LOADING").Msg(resource)
}

// Loaded 已加载
func Loaded(resource string) {
	log.Info().Str("📦", "LOADED").Msg(resource)
}

// ==================== 模块专用方法 ====================

// Database 数据库日志
func Database() *zerolog.Event {
	return log.Info().Str("🗄️", "DATABASE")
}

// Redis Redis日志
func Redis() *zerolog.Event {
	return log.Info().Str("🎯", "REDIS")
}

// HTTP HTTP日志
func HTTP() *zerolog.Event {
	return log.Info().Str("🌐", "HTTP")
}

// Service 服务日志
func Service(name string) *zerolog.Event {
	return log.Info().Str("⚙️", "SERVICE").Str("name", name)
}

// API API日志
func API() *zerolog.Event {
	return log.Info().Str("🔗", "API")
}
