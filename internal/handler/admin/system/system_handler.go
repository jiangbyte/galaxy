package system

import (
	"galaxy/internal/service/admin/system"
	"galaxy/pkg/handler"
	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
	handler.BaseHandler
	systemService *system.SystemService
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{
		systemService: system.NewSystemService(),
	}
}

// GetSysConfigs 获取所有系统配置
func (h *SystemHandler) GetSysConfigs(c *gin.Context) {
	configs, err := h.systemService.GetAllSysConfigs()
	if err != nil {
		h.InternalServerError(c, "获取系统配置失败")
		return
	}

	h.Success(c, gin.H{
		"configs": configs,
	})
}

// UpdateSysConfig 更新系统配置
func (h *SystemHandler) UpdateSysConfig(c *gin.Context) {
	code := c.Param("code")
	var req struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	if err := h.systemService.UpdateSysConfig(code, req.Value); err != nil {
		h.InternalServerError(c, "更新系统配置失败")
		return
	}

	h.SuccessWithMessage(c, nil, "更新成功")
}

// GetSysDict 获取系统字典
func (h *SystemHandler) GetSysDict(c *gin.Context) {
	dictType := c.Param("type")
	dicts, err := h.systemService.GetSysDict(dictType)
	if err != nil {
		h.InternalServerError(c, "获取字典失败")
		return
	}

	h.Success(c, gin.H{
		"dicts": dicts,
	})
}

// GetSysMenus 获取系统菜单
func (h *SystemHandler) GetSysMenus(c *gin.Context) {
	menus, err := h.systemService.GetSysMenus()
	if err != nil {
		h.InternalServerError(c, "获取菜单失败")
		return
	}

	h.Success(c, gin.H{
		"menus": menus,
	})
}
