package model

import (
	"galaxy/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// BaseModel 基础模型，包含所有表的公共字段
type BaseModel struct {
	ID         string     `gorm:"column:id;primaryKey;type:varchar(32)"`
	Deleted    bool       `gorm:"column:deleted;default:false;index"`
	DeleteTime *time.Time `gorm:"column:delete_time"`
	DeleteUser *string    `gorm:"column:delete_user;type:varchar(32)"`
	CreateTime time.Time  `gorm:"column:create_time;default:now()"`
	CreateUser *string    `gorm:"column:create_user;type:varchar(32)"`
	UpdateTime time.Time  `gorm:"column:update_time;default:now()"`
	UpdateUser *string    `gorm:"column:update_user;type:varchar(32)"`
}

// BeforeCreate 在创建记录前生成UUID
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = utils.GenerateID()
	}
	if m.CreateTime.IsZero() {
		m.CreateTime = time.Now()
	}
	if m.UpdateTime.IsZero() {
		m.UpdateTime = time.Now()
	}
	return nil
}

// BeforeUpdate 在更新记录前更新时间
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateTime = time.Now()
	return nil
}

// BeforeDelete 在删除记录前执行软删除
func (m *BaseModel) BeforeDelete(tx *gorm.DB) error {
	// 如果已经是软删除状态，直接返回
	if m.Deleted {
		return nil
	}

	// 执行软删除操作
	now := time.Now()
	updateData := map[string]interface{}{
		"deleted":     true,
		"delete_time": now,
		"update_time": now,
	}

	// 使用 Updates 来更新字段，阻止真正的删除操作
	err := tx.Model(m).Updates(updateData).Error
	if err != nil {
		return err
	}

	// 通过返回错误或修改 tx 来阻止删除
	// 这里修改 Statement 来阻止后续的 DELETE 操作
	tx.Statement.SQL.Reset()
	tx.Statement.SQL.WriteString("-- soft deleted, skip actual delete")
	return nil
}
