package router

import (
	"galaxy/internal/handler/share/config"
	"galaxy/internal/handler/web/user"
	"github.com/gin-gonic/gin"
)

// SetupWebRoutes 设置 Web 业务路由
func SetupWebRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")

	authHandler := user.NewAuthHandler()
	userHandler := user.NewUserHandler()
	configHandler := config.NewConfigHandler()

	// ==================== 公开路由 ====================
	public := api.Group("")
	{
		// 认证相关
		auth := public.Group("/auth")
		{
			auth.GET("/captcha", authHandler.GenerateCaptcha) // 生成验证码
			auth.POST("/register", authHandler.DoRegister)    // 用户注册
			auth.POST("/login", authHandler.DoLogin)          // 用户登录
		}

		// 公开信息
		open := public.Group("/open")
		{
			open.GET("/users/:id", userHandler.GetUserByID) // 获取用户公开信息
		}
	}

	// ==================== 需要认证的路由 ====================
	protected := api.Group("")
	//protected.Use(middleware.AuthMiddleware())
	{
		// 认证相关（需要认证的部分）
		auth := protected.Group("/auth")
		{
			auth.DELETE("/logout", authHandler.DoLogout) // 用户登出（需要认证）
		}

		// 用户管理
		userGroup := protected.Group("/user")
		{
			userGroup.GET("/profile", userHandler.GetUserProfile)
		}

		// 配置管理
		configGroup := protected.Group("/config")
		{
			// 配置分组管理
			group := configGroup.Group("/groups")
			{
				group.POST("", configHandler.CreateConfigGroup)
				group.GET("", configHandler.ConfigGroupList)
				group.GET("/:id", configHandler.GetConfigGroup)
				group.PUT("/:id", configHandler.UpdateConfigGroup)
				group.DELETE("/:id", configHandler.DeleteConfigGroup)
			}

			// 配置项管理
			item := configGroup.Group("/items")
			{
				item.POST("", configHandler.CreateConfigItem)
				item.GET("", configHandler.ConfigItemList)
				item.GET("/:id", configHandler.GetConfigItem)
				item.PUT("/:id", configHandler.UpdateConfigItem)
				item.DELETE("/:id", configHandler.DeleteConfigItem)
			}
		}
	}
}
