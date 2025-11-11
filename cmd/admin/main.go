// cmd/admin/main.go
package main

import (
	"fmt"
	"galaxy/pkg/config"
	"galaxy/pkg/database"
	"galaxy/pkg/logger"
	"galaxy/pkg/redis"
	"galaxy/pkg/router"
	"github.com/gin-gonic/gin"
	"os"
)

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

	// 创建路由
	r := router.NewRouter()
	r.SetupAdminRoutes()

	// 启动服务
	port := fmt.Sprintf(":%d", cfg.Server.Admin.Port)
	logger.Service("Admin").Str("port", fmt.Sprintf("%d", cfg.Server.Admin.Port))
	if err := r.GetEngine().Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Admin Error")
	}
}
