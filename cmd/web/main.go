package main

import (
	"fmt"
	"galaxy/internal/router"
	"galaxy/pkg/captcha"
	"galaxy/pkg/config"
	"galaxy/pkg/database"
	"galaxy/pkg/logger"
	"galaxy/pkg/redis"
	pkgRouter "galaxy/pkg/router"
	"github.com/gin-gonic/gin"
	"os"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    "ok",
		"service":   "galaxy-web",
		"timestamp": gin.H{},
	})
}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info().Msg("🚀 Gin running in RELEASE mode")
	} else {
		logger.Info().Msg("🔧 Gin running in DEBUG mode")
	}

	// 加载配置
	cfg := config.Load("configs/config.yaml")

	// 初始化数据库
	database.Init()

	// 初始化Redis
	redis.Init()

	// 初始化验证码
	captcha.Init()

	// 创建路由引擎
	engine := pkgRouter.NewEngine()

	// 设置业务路由
	router.SetupWebRoutes(engine)

	// 打印路由信息
	pkgRouter.PrintRoutes(engine)

	// 启动服务
	port := fmt.Sprintf(":%d", cfg.Server.Web.Port)

	logger.Info().
		Msg(fmt.Sprintf("🚀 Web service starts on port: %s", port))

	if err := engine.Run(port); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Web service startup failed")
	}
}
