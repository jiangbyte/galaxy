package models

import (
	"galaxy/pkg/model"
	"time"
)

// ============================================================
// 头衔体系
// ============================================================

// TitleConfig 头衔配置表
type TitleConfig struct {
	model.BaseModel
	Title       string  `gorm:"column:title;type:varchar(100);not null;uniqueIndex:idx_title_name"` // 头衔名称
	Code        string  `gorm:"column:code;type:varchar(50);not null;uniqueIndex:idx_title_code"`   // 头衔编码
	Description string  `gorm:"column:description;type:varchar(500);not null"`                      // 头衔描述
	Color       string  `gorm:"column:color;type:varchar(20);default:'#666666'"`                    // 头衔颜色
	Rarity      int     `gorm:"column:rarity;type:smallint;default:1"`                              // 稀有度 1-5
	Category    string  `gorm:"column:category;type:varchar(50);default:'achievement'"`             // 分类: achievement/vip/level/honor/special
	Icon        *string `gorm:"column:icon;type:varchar(255)"`                                      // 头衔图标
	Background  *string `gorm:"column:background;type:varchar(255)"`                                // 头衔背景
	// 获取条件
	ConditionType  string `gorm:"column:condition_type;type:varchar(50);not null"`   // 条件类型: level/vip/achievement/duration
	ConditionValue string `gorm:"column:condition_value;type:varchar(255);not null"` // 条件值
	// 有效期设置
	IsPermanent  bool       `gorm:"column:is_permanent;default:true"` // 是否永久头衔
	DurationDays int        `gorm:"column:duration_days;default:0"`   // 有效期天数(0表示永久)
	IsLimited    bool       `gorm:"column:is_limited;default:false"`  // 是否限时头衔
	StartTime    *time.Time `gorm:"column:start_time"`                // 开始时间(限时头衔)
	EndTime      *time.Time `gorm:"column:end_time"`                  // 结束时间(限时头衔)
	Weight       int        `gorm:"column:weight;default:0"`          // 权重，用于排序
}

func (TitleConfig) TableName() string {
	return "title_config"
}

// UserTitle 用户头衔表
type UserTitle struct {
	model.BaseModel
	AccountID   string     `gorm:"column:account_id;type:varchar(32);not null;index:idx_user_title_account"`
	TitleID     string     `gorm:"column:title_id;type:varchar(32);not null"`                 // 头衔ID
	AcquireTime time.Time  `gorm:"column:acquire_time;default:now()"`                         // 获取时间
	IsEquipped  bool       `gorm:"column:is_equipped;default:false;index:idx_title_equipped"` // 是否装备
	Source      string     `gorm:"column:source;type:varchar(50);default:'system'"`           // 获取来源
	ExpireTime  *time.Time `gorm:"column:expire_time;index:idx_title_expire"`                 // 过期时间
}

func (UserTitle) TableName() string {
	return "user_title"
}
