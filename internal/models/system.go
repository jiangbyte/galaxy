package models

import (
	"galaxy/pkg/model"
	"time"

	"gorm.io/datatypes"
)

// ============================================================
// 系统基础表
// ============================================================

// SysConfig 系统配置表
type SysConfig struct {
	model.BaseModel
	Name          string  `gorm:"column:name;type:varchar(255)"`
	Code          string  `gorm:"column:code;type:varchar(255);uniqueIndex:idx_code"`
	Value         string  `gorm:"column:value;type:varchar(255)"`
	ComponentType *string `gorm:"column:component_type;type:varchar(255)"`
	Description   *string `gorm:"column:description;type:varchar(255)"`
	ConfigType    *string `gorm:"column:config_type;type:varchar(255)"`
	Sort          int     `gorm:"column:sort;default:0"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}

// SysDict 系统字典表
type SysDict struct {
	model.BaseModel
	DictType  string  `gorm:"column:dict_type;type:varchar(64);uniqueIndex:uk_type_code"`
	TypeLabel *string `gorm:"column:type_label;type:varchar(64)"`
	DictValue string  `gorm:"column:dict_value;type:varchar(255);uniqueIndex:uk_type_code"`
	DictLabel *string `gorm:"column:dict_label;type:varchar(255)"`
	SortOrder int     `gorm:"column:sort_order;default:0"`
}

func (SysDict) TableName() string {
	return "sys_dict"
}

// SysLog 系统活动/日志记录表
type SysLog struct {
	model.BaseModel
	UserID        *string    `gorm:"column:user_id;type:varchar(32)"`
	Operation     *string    `gorm:"column:operation;type:varchar(255)"`
	Method        *string    `gorm:"column:method;type:varchar(255)"`
	Params        *string    `gorm:"column:params;type:text"`
	IP            *string    `gorm:"column:ip;type:varchar(255)"`
	OperationTime *time.Time `gorm:"column:operation_time"`
	Category      *string    `gorm:"column:category;type:varchar(255)"`
	Module        *string    `gorm:"column:module;type:varchar(255)"`
	Description   *string    `gorm:"column:description;type:varchar(255)"`
	Status        *string    `gorm:"column:status;type:varchar(255)"`
	Message       *string    `gorm:"column:message;type:text"`
}

func (SysLog) TableName() string {
	return "sys_log"
}

// SysMenu 菜单表
type SysMenu struct {
	model.BaseModel
	PID           *string        `gorm:"column:pid;type:varchar(32);default:0;index:idx_pid"`
	Name          *string        `gorm:"column:name;type:varchar(100)"`
	Path          *string        `gorm:"column:path;type:varchar(200)"`
	ComponentPath *string        `gorm:"column:component_path;type:varchar(500)"`
	Title         *string        `gorm:"column:title;type:varchar(100)"`
	Icon          *string        `gorm:"column:icon;type:varchar(100)"`
	KeepAlive     bool           `gorm:"column:keep_alive;default:false"`
	Visible       bool           `gorm:"column:visible;default:true"`
	Sort          int            `gorm:"column:sort;default:0;index:idx_sort"`
	Pined         bool           `gorm:"column:pined;default:false"`
	MenuType      int            `gorm:"column:menu_type;default:0;index:idx_menu_type"`
	Parameters    *string        `gorm:"column:parameters;type:varchar(500)"`
	ExtraParams   datatypes.JSON `gorm:"column:extra_params;type:jsonb"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
