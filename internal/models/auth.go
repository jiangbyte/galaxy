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
	GroupID   *string `gorm:"column:group_id;type:varchar(32)"` // 账户所属组
	// 账户安全状态
	Status int `gorm:"column:status;type:smallint;default:0"` // 0: 正常 1: 锁定 2: 禁用
	// 安全相关
	PasswordStrength   int        `gorm:"column:password_strength;type:smallint;default:0"` // 密码强度 0 - 3
	LastPasswordChange *time.Time `gorm:"column:last_password_change"`                      // 最后修改密码的时间
	// 登录与活跃信息
	LastLoginTime *time.Time `gorm:"column:last_login_time"`                // 最后登录时间
	LastLoginIP   *string    `gorm:"column:last_login_ip;type:varchar(64)"` // 最后登录IP
	LoginCount    int        `gorm:"column:login_count;default:0"`          // 登录次数
}

// LoginInfoUpdate 可以定义一个专门的更新结构体
type LoginInfoUpdate struct {
	LastLoginTime *time.Time `gorm:"column:last_login_time"`
	LastLoginIP   *string    `gorm:"column:last_login_ip"`
	LockUntil     *time.Time `gorm:"column:lock_until"`
}

func (AuthAccount) TableName() string {
	return "auth_account"
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
