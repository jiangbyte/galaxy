package models

import (
	"galaxy/pkg/model"
	"time"
)

// ============================================================
// VIP 体系
// ============================================================

// VipInfo 用户VIP信息表
type VipInfo struct {
	model.BaseModel
	AccountID     string     `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_vip_account"`
	VIPLevel      int        `gorm:"column:vip_level;type:smallint;default:0;index:idx_vip_level"` // VIP等级 (0:普通用户, 1-10:VIP等级)
	VIPExpireTime *time.Time `gorm:"column:vip_expire_time;index:idx_vip_expire"`                  // VIP到期时间
	VIPStartTime  *time.Time `gorm:"column:vip_start_time"`                                        // VIP开始时间
	IsAutoRenew   bool       `gorm:"column:is_auto_renew;default:false"`                           // 是否自动续费
	TotalMonths   int        `gorm:"column:total_months;default:0"`                                // 累计开通月数
	UpgradeTime   *time.Time `gorm:"column:upgrade_time"`                                          // 最后升级时间
}

func (VipInfo) TableName() string {
	return "vip_info"
}

// VipPrivilege VIP等级权益配置表
type VipPrivilege struct {
	model.BaseModel
	VIPLevel     int     `gorm:"column:vip_level;type:smallint;not null;uniqueIndex:idx_vip_privilege_level"` // VIP等级
	Name         string  `gorm:"column:name;type:varchar(100);not null"`                                      // 权益名称
	Description  string  `gorm:"column:description;type:varchar(500)"`                                        // 权益描述
	Icon         *string `gorm:"column:icon;type:varchar(255)"`                                               // 权益图标
	Color        string  `gorm:"column:color;type:varchar(20);default:'#FFD700'"`                             // 权益颜色
	Weight       int     `gorm:"column:weight;default:0"`                                                     // 权重，用于排序
	PrivilegeKey string  `gorm:"column:privilege_key;type:varchar(50);not null"`                              // 权益键名，用于代码识别
}

func (VipPrivilege) TableName() string {
	return "vip_privilege"
}

// VipLevelConfig VIP等级配置表
type VipLevelConfig struct {
	model.BaseModel
	VIPLevel    int     `gorm:"column:vip_level;type:smallint;not null;uniqueIndex:idx_vip_level_config"` // VIP等级
	LevelName   string  `gorm:"column:level_name;type:varchar(100);not null"`                             // 等级名称
	MonthPrice  float64 `gorm:"column:month_price;type:decimal(10,2);not null"`                           // 月价格
	YearPrice   float64 `gorm:"column:year_price;type:decimal(10,2)"`                                     // 年价格
	ExpRequired int64   `gorm:"column:exp_required;default:0"`                                            // 升级所需经验/积分
	BadgeID     *string `gorm:"column:badge_id;type:varchar(32)"`                                         // 对应徽章ID
	TitleID     *string `gorm:"column:title_id;type:varchar(32)"`                                         // 对应头衔ID
	Color       string  `gorm:"column:color;type:varchar(20);default:'#FFD700'"`                          // 等级颜色
	MaxLevel    bool    `gorm:"column:max_level;default:false"`                                           // 是否最高等级
}

func (VipLevelConfig) TableName() string {
	return "vip_level_config"
}
