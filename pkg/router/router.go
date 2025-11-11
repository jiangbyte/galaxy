package router

import (
	"galaxy/pkg/handler"
	"galaxy/pkg/logger"
	"galaxy/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// NewEngine 创建配置好的 Gin 引擎（只包含基础中间件）
func NewEngine() *gin.Engine {
	engine := gin.New()

	// 基础中间件
	engine.Use(logger.GinLogger())
	engine.Use(logger.GinRecovery())
	engine.Use(middleware.CORS())

	// 404 使用
	engine.NoRoute(handler.NotFoundHandler())

	return engine
}

// PrintRoutes 打印路由信息
func PrintRoutes(engine *gin.Engine) {
	logger.Info().Msg("📋 Registered Routes:")

	for _, route := range engine.Routes() {
		var methodEmoji string
		switch route.Method {
		case "GET":
			methodEmoji = "💙"
		case "POST":
			methodEmoji = "💚"
		case "PUT":
			methodEmoji = "🟡"
		case "DELETE":
			methodEmoji = "❤️"
		case "PATCH":
			methodEmoji = "🟠"
		default:
			methodEmoji = "⚪"
		}

		logger.Info().
			Str("path", route.Path).
			Str("method", methodEmoji+" "+route.Method).
			Msg("Route")
	}
}
