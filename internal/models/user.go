package models

import (
	"galaxy/pkg/model"
	"time"

	"gorm.io/datatypes"
)

// ============================================================
// 用户认证
// ============================================================

// AuthAccount 核心账户表
type AuthAccount struct {
	model.BaseModel
	Username  string  `gorm:"column:username;type:varchar(64);not null;uniqueIndex:idx_account_username"`
	Password  string  `gorm:"column:password;type:varchar(100);not null"`
	Email     *string `gorm:"column:email;type:varchar(128);uniqueIndex:idx_account_email"`
	Telephone *string `gorm:"column:telephone;type:varchar(20);uniqueIndex:idx_account_telephone"`
	GroupID   *string `gorm:"column:group_id;type:varchar(32)"`

	// 账户安全状态
	IsVerified bool `gorm:"column:is_verified;default:false"`
	IsActive   bool `gorm:"column:is_active;default:true"`

	// 安全相关
	PasswordStrength   int        `gorm:"column:password_strength;type:smallint;default:0"`
	LastPasswordChange *time.Time `gorm:"column:last_password_change"`
	SecurityQuestion   *string    `gorm:"column:security_question;type:varchar(255)"`
	SecurityAnswer     *string    `gorm:"column:security_answer;type:varchar(255)"`

	// 登录与活跃信息
	LastLoginTime    *time.Time `gorm:"column:last_login_time"`
	LastLoginIP      *string    `gorm:"column:last_login_ip;type:varchar(64)"`
	LastActiveTime   *time.Time `gorm:"column:last_active_time"`
	LoginCount       int        `gorm:"column:login_count;default:0"`
	FailedLoginCount int        `gorm:"column:failed_login_count;type:smallint;default:0"`
	LockUntil        *time.Time `gorm:"column:lock_until"`
}

func (AuthAccount) TableName() string {
	return "auth_account"
}

// AuthUserInfo 用户基本信息表
type AuthUserInfo struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`

	// 基础身份信息
	Nickname string     `gorm:"column:nickname;type:varchar(128);not null;index:idx_user_nickname"`
	RealName *string    `gorm:"column:real_name;type:varchar(64)"`
	Avatar   *string    `gorm:"column:avatar;type:varchar(255)"`
	Gender   int        `gorm:"column:gender;type:smallint;default:0"`
	Birthday *time.Time `gorm:"column:birthday;type:date"`

	// 展示信息
	DisplayName *string `gorm:"column:display_name;type:varchar(128)"`
	Title       *string `gorm:"column:title;type:varchar(100)"`
	Signature   *string `gorm:"column:signature;type:varchar(500)"`

	// 基础统计
	Level int   `gorm:"column:level;default:1;index:idx_user_level"`
	Exp   int64 `gorm:"column:exp;default:0"`
}

func (AuthUserInfo) TableName() string {
	return "auth_user_info"
}

// AuthUserProfile 用户档案详情表
type AuthUserProfile struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`

	// 教育职业信息
	School    *string `gorm:"column:school;type:varchar(100)"`
	Major     *string `gorm:"column:major;type:varchar(100)"`
	StudentID *string `gorm:"column:student_id;type:varchar(50)"`
	Company   *string `gorm:"column:company;type:varchar(100)"`
	JobTitle  *string `gorm:"column:job_title;type:varchar(100)"`
	Industry  *string `gorm:"column:industry;type:varchar(100)"`

	// 地理位置
	Country  *string `gorm:"column:country;type:varchar(50)"`
	Province *string `gorm:"column:province;type:varchar(50)"`
	City     *string `gorm:"column:city;type:varchar(50)"`
	Location *string `gorm:"column:location;type:varchar(100)"`

	// 个人背景
	Background *string        `gorm:"column:background;type:varchar(255)"`
	Bio        *string        `gorm:"column:bio;type:text"`
	Interests  datatypes.JSON `gorm:"column:interests;type:jsonb"`

	// 社交链接
	Website *string `gorm:"column:website;type:varchar(255)"`
	GitHub  *string `gorm:"column:github;type:varchar(100)"`
	Blog    *string `gorm:"column:blog;type:varchar(255)"`
	Weibo   *string `gorm:"column:weibo;type:varchar(100)"`

	// 隐私设置
	PrivacyLevel int  `gorm:"column:privacy_level;type:smallint;default:1"`
	ShowRealName bool `gorm:"column:show_real_name;default:false"`
	ShowBirthday bool `gorm:"column:show_birthday;default:false"`
	ShowLocation bool `gorm:"column:show_location;default:true"`
}

func (AuthUserProfile) TableName() string {
	return "auth_user_profile"
}

// AuthUserPreference 用户偏好设置表
type AuthUserPreference struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`

	// 界面设置
	Theme     string `gorm:"column:theme;type:varchar(50);default:light"`
	Language  string `gorm:"column:language;type:varchar(10);default:zh-CN"`
	FontSize  int    `gorm:"column:font_size;type:smallint;default:14"`
	CodeTheme string `gorm:"column:code_theme;type:varchar(50);default:github"`

	// 通知设置
	EmailNotifications bool `gorm:"column:email_notifications;default:true"`
	PushNotifications  bool `gorm:"column:push_notifications;default:true"`

	// 隐私与展示
	AllowDirectMessage bool `gorm:"column:allow_direct_message;default:true"`
}

func (AuthUserPreference) TableName() string {
	return "auth_user_preference"
}

// AuthUserStats 用户统计信息表
type AuthUserStats struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`

	// 等级与经验
	Level int     `gorm:"column:level;default:1;index:idx_user_level"`
	Exp   int64   `gorm:"column:exp;default:0"`
	Title *string `gorm:"column:title;type:varchar(100)"`
}

func (AuthUserStats) TableName() string {
	return "auth_user_stats"
}

// AuthUserVIP VIP信息表
type AuthUserVIP struct {
	model.BaseModel
	AccountID     string     `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	VIPLevel      int        `gorm:"column:vip_level;type:smallint;default:0"`
	VIPExpireTime *time.Time `gorm:"column:vip_expire_time;index:idx_vip_expire"`
	VIPStartTime  *time.Time `gorm:"column:vip_start_time"`
	IsAutoRenew   bool       `gorm:"column:is_auto_renew;default:false"`
}

func (AuthUserVIP) TableName() string {
	return "auth_user_vip"
}

// AuthUserBadge 用户徽章表
type AuthUserBadge struct {
	model.BaseModel
	AccountID   string    `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	BadgeID     string    `gorm:"column:badge_id;type:varchar(32);not null"`
	BadgeName   *string   `gorm:"column:badge_name;type:varchar(100)"`
	BadgeIcon   *string   `gorm:"column:badge_icon;type:varchar(255)"`
	AcquireTime time.Time `gorm:"column:acquire_time;default:now()"`
	IsEquipped  bool      `gorm:"column:is_equipped;default:false;index:idx_badge_equipped"`
}

func (AuthUserBadge) TableName() string {
	return "auth_user_badge"
}

// AuthBadgeConfig 徽章配置表
type AuthBadgeConfig struct {
	model.BaseModel
	Name           string  `gorm:"column:name;type:varchar(100);not null"`
	Code           string  `gorm:"column:code;type:varchar(50);not null;uniqueIndex:idx_badge_code"`
	Icon           *string `gorm:"column:icon;type:varchar(255)"`
	Description    *string `gorm:"column:description;type:varchar(500)"`
	Color          *string `gorm:"column:color;type:varchar(20)"`
	Rarity         int     `gorm:"column:rarity;type:smallint;default:1"`
	ConditionType  *string `gorm:"column:condition_type;type:varchar(50)"`
	ConditionValue *string `gorm:"column:condition_value;type:varchar(255)"`
}

func (AuthBadgeConfig) TableName() string {
	return "auth_badge_config"
}

// AuthUserVIPPrivilege VIP等级权益表
type AuthUserVIPPrivilege struct {
	model.BaseModel
	VIPLevel        int     `gorm:"column:vip_level;type:smallint"`
	Name            *string `gorm:"column:name;type:varchar(50)"`
	Color           *string `gorm:"column:color;type:varchar(20)"`
	AIChatPrivilege *bool   `gorm:"column:ai_chat_privilege"`
}

func (AuthUserVIPPrivilege) TableName() string {
	return "auth_user_vip_privilege"
}

// AuthUserLevelConfig 用户等级成长配置表
type AuthUserLevelConfig struct {
	model.BaseModel
	Level         int     `gorm:"column:level;not null;uniqueIndex:idx_level_config_level"`
	LevelName     *string `gorm:"column:level_name;type:varchar(100)"`
	ExpRequired   int64   `gorm:"column:exp_required"`
	BadgeUnlocked *string `gorm:"column:badge_unlocked;type:varchar(100)"`
	TitleUnlocked *string `gorm:"column:title_unlocked;type:varchar(100)"`
}

func (AuthUserLevelConfig) TableName() string {
	return "auth_user_level_config"
}

// AuthAccountRole 账户-角色 关联表
type AuthAccountRole struct {
	ID        string `gorm:"column:id;primaryKey;type:varchar(32)"`
	AccountID string `gorm:"column:account_id;type:varchar(32);not null"`
	RoleID    string `gorm:"column:role_id;type:varchar(32);not null"`
}

func (AuthAccountRole) TableName() string {
	return "auth_account_role"
}

// AuthGroup 用户组表
type AuthGroup struct {
	model.BaseModel
	ParentID     *string `gorm:"column:parent_id;type:varchar(32)"`
	Name         *string `gorm:"column:name;type:varchar(100)"`
	Code         *string `gorm:"column:code;type:varchar(50);uniqueIndex:idx_group_code"`
	Description  *string `gorm:"column:description;type:varchar(255)"`
	Sort         int     `gorm:"column:sort;type:smallint;default:99"`
	AdminID      *string `gorm:"column:admin_id;type:varchar(32)"`
	MaxUserCount *int    `gorm:"column:max_user_count"`
	IsSystem     bool    `gorm:"column:is_system;default:false"`
}

func (AuthGroup) TableName() string {
	return "auth_group"
}

// AuthRole 角色表
type AuthRole struct {
	model.BaseModel
	Name           *string        `gorm:"column:name;type:varchar(255);index:idx_name"`
	Code           *string        `gorm:"column:code;type:varchar(50);index:idx_role_code"`
	DataScope      *string        `gorm:"column:data_scope;type:varchar(50);index:idx_data_scope"`
	Description    *string        `gorm:"column:description;type:varchar(255)"`
	AssignGroupIDs datatypes.JSON `gorm:"column:assign_group_ids;type:jsonb"`
}

func (AuthRole) TableName() string {
	return "auth_role"
}

// AuthRoleMenu 角色-菜单 关联表
type AuthRoleMenu struct {
	ID     string `gorm:"column:id;primaryKey;type:varchar(32)"`
	RoleID string `gorm:"column:role_id;type:varchar(32);not null"`
	MenuID string `gorm:"column:menu_id;type:varchar(32);not null"`
}

func (AuthRoleMenu) TableName() string {
	return "auth_role_menu"
}
