package config

import (
	"galaxy/internal/models"
	config2 "galaxy/internal/query/config"
	"galaxy/internal/service/share/config"
	"galaxy/pkg/handler"

	"github.com/gin-gonic/gin"
)

type ConfigHandler struct {
	handler.BaseHandler
	configService config.ConfigService
}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{
		configService: config.NewConfigService(),
	}
}

// ============================================================
// ConfigGroup Handlers
// ============================================================

// CreateConfigGroup 创建配置分组
func (h *ConfigHandler) CreateConfigGroup(c *gin.Context) {
	h.StartTimer(c)

	var group models.ConfigGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	if err := h.configService.CreateConfigGroup(&group); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, group)
}

// GetConfigGroup 获取配置分组详情
func (h *ConfigHandler) GetConfigGroup(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "分组ID不能为空")
		return
	}

	group, err := h.configService.GetConfigGroup(id)
	if err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, group)
}

// UpdateConfigGroup 更新配置分组
func (h *ConfigHandler) UpdateConfigGroup(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "分组ID不能为空")
		return
	}

	var group models.ConfigGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	group.ID = id
	if err := h.configService.UpdateConfigGroup(&group); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, group)
}

// DeleteConfigGroup 删除配置分组
func (h *ConfigHandler) DeleteConfigGroup(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "分组ID不能为空")
		return
	}

	if err := h.configService.DeleteConfigGroup(id); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, nil)
}

// ConfigGroupList 获取配置分组列表
func (h *ConfigHandler) ConfigGroupList(c *gin.Context) {
	h.StartTimer(c)

	var req config2.ConfigGroupQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	result, err := h.configService.ConfigGroupList(&req)
	if err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, result)
}

// ============================================================
// ConfigItem Handlers
// ============================================================

// CreateConfigItem 创建配置项
func (h *ConfigHandler) CreateConfigItem(c *gin.Context) {
	h.StartTimer(c)

	var item models.ConfigItem
	if err := c.ShouldBindJSON(&item); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	if err := h.configService.CreateConfigItem(&item); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, item)
}

// GetConfigItem 获取配置项详情
func (h *ConfigHandler) GetConfigItem(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "配置项ID不能为空")
		return
	}

	item, err := h.configService.GetConfigItem(id)
	if err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, item)
}

// UpdateConfigItem 更新配置项
func (h *ConfigHandler) UpdateConfigItem(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "配置项ID不能为空")
		return
	}

	var item models.ConfigItem
	if err := c.ShouldBindJSON(&item); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	item.ID = id
	if err := h.configService.UpdateConfigItem(&item); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, item)
}

// DeleteConfigItem 删除配置项
func (h *ConfigHandler) DeleteConfigItem(c *gin.Context) {
	h.StartTimer(c)

	id := c.Param("id")
	if id == "" {
		h.BadRequest(c, "配置项ID不能为空")
		return
	}

	if err := h.configService.DeleteConfigItem(id); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, nil)
}

// ConfigItemList 获取配置项列表
func (h *ConfigHandler) ConfigItemList(c *gin.Context) {
	h.StartTimer(c)

	var req config2.ConfigItemQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	result, err := h.configService.ConfigItemList(&req)
	if err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, result)
}
