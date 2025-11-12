package models

import (
	"galaxy/pkg/model"
	"time"
)

// ============================================================
// 徽章体系
// ============================================================

// BadgeConfig 徽章配置表
type BadgeConfig struct {
	model.BaseModel
	Name           string     `gorm:"column:name;type:varchar(100);not null"`                           // 徽章名称
	Code           string     `gorm:"column:code;type:varchar(50);not null;uniqueIndex:idx_badge_code"` // 徽章编码
	Icon           string     `gorm:"column:icon;type:varchar(255);not null"`                           // 徽章图标
	Description    string     `gorm:"column:description;type:varchar(500);not null"`                    // 徽章描述
	Color          string     `gorm:"column:color;type:varchar(20);default:'#666666'"`                  // 徽章颜色
	Rarity         int        `gorm:"column:rarity;type:smallint;default:1"`                            // 稀有度 1-5 (1:普通, 2:稀有, 3:史诗, 4:传说, 5:限定)
	Category       string     `gorm:"column:category;type:varchar(50);default:'achievement'"`           // 分类: achievement/vip/level/event/special
	ConditionType  string     `gorm:"column:condition_type;type:varchar(50);not null"`                  // 条件类型: level/vip_expire/daily_login/content_count...
	ConditionValue string     `gorm:"column:condition_value;type:varchar(255);not null"`                // 条件值
	IsHidden       bool       `gorm:"column:is_hidden;default:false"`                                   // 是否隐藏徽章
	IsLimited      bool       `gorm:"column:is_limited;default:false"`                                  // 是否限时徽章
	StartTime      *time.Time `gorm:"column:start_time"`                                                // 开始时间(限时徽章)
	EndTime        *time.Time `gorm:"column:end_time"`                                                  // 结束时间(限时徽章)
	Weight         int        `gorm:"column:weight;default:0"`                                          // 权重，用于排序
}

func (BadgeConfig) TableName() string {
	return "badge_config"
}

// UserBadge 用户徽章表
type UserBadge struct {
	model.BaseModel
	AccountID   string     `gorm:"column:account_id;type:varchar(32);not null;index:idx_user_badge_account"`
	BadgeID     string     `gorm:"column:badge_id;type:varchar(32);not null"`                 // 徽章ID
	AcquireTime time.Time  `gorm:"column:acquire_time;default:now()"`                         // 获取时间
	IsEquipped  bool       `gorm:"column:is_equipped;default:false;index:idx_badge_equipped"` // 是否已装备
	Source      string     `gorm:"column:source;type:varchar(50);default:'system'"`           // 获取来源: system/activity/purchase/gift
	ExpireTime  *time.Time `gorm:"column:expire_time;index:idx_badge_expire"`                 // 过期时间(临时徽章)
}

func (UserBadge) TableName() string {
	return "user_badge"
}
