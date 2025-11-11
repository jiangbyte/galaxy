package router

import (
	"galaxy/internal/handler/web/user"
	"galaxy/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// SetupWebRoutes 设置 Web 业务路由
func SetupWebRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")

	// ==================== 公开路由 ====================
	public := api.Group("")
	{
		// 认证相关
		auth := public.Group("/auth")
		{
			authHandler := user.NewAuthHandler()
			auth.GET("/captcha", authHandler.GenerateCaptcha) // 生成验证码
			auth.POST("/register", authHandler.DoRegister)    // 用户注册
			auth.POST("/login", authHandler.DoLogin)          // 用户登录
		}

		// 公开信息
		open := public.Group("/open")
		{
			userHandler := user.NewUserHandler()
			open.GET("/users/:id", userHandler.GetUserByID) // 获取用户公开信息
		}
	}

	// ==================== 需要认证的路由 ====================
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户管理
		userGroup := protected.Group("/user")
		{
			userHandler := user.NewUserHandler()
			userGroup.GET("/profile", userHandler.GetUserProfile)
		}
	}
}
