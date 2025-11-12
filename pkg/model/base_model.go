package model

import (
	"galaxy/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// BaseModel 基础模型，包含所有表的公共字段
type BaseModel struct {
	ID         string         `gorm:"column:id;primaryKey;type:varchar(32)"`
	Deleted    bool           `gorm:"column:deleted;default:false;index"`
	DeletedAt  gorm.DeletedAt `gorm:"column:delete_time"`
	DeleteUser *string        `gorm:"column:delete_user;type:varchar(32)"`
	CreatedAt  time.Time      `gorm:"column:create_time;default:now()"`
	CreateUser *string        `gorm:"column:create_user;type:varchar(32)"`
	UpdatedAt  time.Time      `gorm:"column:update_time;default:now()"`
	UpdateUser *string        `gorm:"column:update_user;type:varchar(32)"`
}

// BeforeCreate 在创建记录前生成UUID
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = utils.GenerateID()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate 在更新记录前更新时间
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()

	// 如果 DeletedAt 被设置，同步到 Deleted 字段
	if m.DeletedAt.Valid && !m.Deleted {
		m.Deleted = true
	}
	// 如果 DeletedAt 被清空，同步到 Deleted 字段
	if !m.DeletedAt.Valid && m.Deleted {
		m.Deleted = false
	}
	return nil
}
