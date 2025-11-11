package gin

import (
	"fmt"
	"galaxy/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

// Setup 初始化 Gin 引擎
func Setup() *gin.Engine {
	// 设置 Gin 模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info().Msg("🚀 Gin running in RELEASE mode")
	} else {
		gin.SetMode(gin.DebugMode)
		logger.Info().Msg("🔧 Gin running in DEBUG mode")
	}

	// 创建 Gin 实例
	r := gin.New()

	// 使用自定义日志和恢复中间件
	r.Use(logger.GinLogger())
	r.Use(logger.GinRecovery())

	// 设置信任的代理（根据你的部署环境调整）
	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	return r
}

// PrintRoutes 打印路由信息（替代 Gin 的默认调试输出）
func PrintRoutes(r *gin.Engine) {
	logger.Info().Msg("📋 Registered Routes:")

	for _, route := range r.Routes() {
		event := logger.Info()
		switch route.Method {
		case "GET":
			event = event.Str("💙", "GET")
		case "POST":
			event = event.Str("💚", "POST")
		case "PUT":
			event = event.Str("🟡", "PUT")
		case "DELETE":
			event = event.Str("❤️", "DELETE")
		case "PATCH":
			event = event.Str("🟠", "PATCH")
		default:
			event = event.Str("⚪", route.Method)
		}

		event.
			Str("path", route.Path).
			Str("handler", route.Handler).
			Msg("Route registered")
	}
}

// StartServer 启动 HTTP 服务器
func StartServer(r *gin.Engine, port string) {
	addr := ":" + port
	logger.HTTP().
		Str("port", port).
		Msg("Server starting")

	logger.Info().Msg(fmt.Sprintf("🚀 Web服务启动在端口: %s", port))

	if err := r.Run(addr); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Server failed to start")
	}
}
