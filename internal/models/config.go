package models

import "galaxy/pkg/model"

// ============================================================
// 配置表
// ============================================================

// ConfigGroup 配置分组表
type ConfigGroup struct {
	model.BaseModel
	Name        string `gorm:"column:name;type:varchar(100);not null;comment:分组名称"`
	Code        string `gorm:"column:code;type:varchar(100);uniqueIndex;not null;comment:分组代码"`
	Description string `gorm:"column:description;type:varchar(500);comment:分组描述"`
	Sort        int    `gorm:"column:sort;default:0;comment:排序"`
	IsSystem    bool   `gorm:"column:is_system;default:false;comment:是否系统分组"`
}

func (ConfigGroup) TableName() string {
	return "config_group"
}

// ConfigItem 系统配置表
type ConfigItem struct {
	model.BaseModel
	GroupID       string  `gorm:"column:group_id;type:varchar(32);not null;index:idx_group_id;comment:分组ID"`
	Name          string  `gorm:"column:name;type:varchar(255)"`
	Code          string  `gorm:"column:code;type:varchar(255);uniqueIndex:idx_code"`
	Value         string  `gorm:"column:value;type:varchar(255)"`
	ComponentType *string `gorm:"column:component_type;type:varchar(255)"`
	Description   *string `gorm:"column:description;type:varchar(255)"`
	Sort          int     `gorm:"column:sort;default:0"`
}

func (ConfigItem) TableName() string {
	return "config_item"
}
