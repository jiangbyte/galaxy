package router

import (
	"galaxy/internal/handler/admin/system"
	"github.com/gin-gonic/gin"
)

// SetupAdminRoutes 设置管理后台路由
func SetupAdminRoutes(engine *gin.Engine) {
	adminGroup := engine.Group("/admin/api/v1")
	systemHandler := system.NewSystemHandler()

	// 系统管理路由
	systemGroup := adminGroup.Group("/system")
	{
		systemGroup.GET("/configs", systemHandler.GetSysConfigs)
		systemGroup.PUT("/configs/:code", systemHandler.UpdateSysConfig)
		systemGroup.GET("/dict/:type", systemHandler.GetSysDict)
		systemGroup.GET("/menus", systemHandler.GetSysMenus)
	}
}
