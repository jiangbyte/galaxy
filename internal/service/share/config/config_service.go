package config

import (
	"errors"
	"galaxy/pkg/database"

	"galaxy/internal/models"
	"galaxy/internal/query/config"
	"galaxy/pkg/query"
	"gorm.io/gorm"
)

type ConfigService interface {
	// ConfigGroup CRUD
	CreateConfigGroup(group *models.ConfigGroup) error
	GetConfigGroup(id string) (*models.ConfigGroup, error)
	UpdateConfigGroup(group *models.ConfigGroup) error
	DeleteConfigGroup(id string) error
	ConfigGroupList(req *config.ConfigGroupQueryRequest) (*query.PaginationResponse[models.ConfigGroup], error)

	// ConfigItem CRUD
	CreateConfigItem(item *models.ConfigItem) error
	GetConfigItem(id string) (*models.ConfigItem, error)
	UpdateConfigItem(item *models.ConfigItem) error
	DeleteConfigItem(id string) error
	ConfigItemList(req *config.ConfigItemQueryRequest) (*query.PaginationResponse[models.ConfigItem], error)
}

type ConfigServiceImpl struct {
	db *gorm.DB
}

func NewConfigService() ConfigService {
	return &ConfigServiceImpl{
		db: database.GetDB(),
	}
}

// ============================================================
// ConfigGroup CRUD 实现
// ============================================================

// CreateConfigGroup 创建配置分组
func (s *ConfigServiceImpl) CreateConfigGroup(group *models.ConfigGroup) error {
	if group == nil {
		return errors.New("配置组不能为 nil")
	}

	// 检查分组代码是否已存在
	var count int64
	if err := s.db.Model(&models.ConfigGroup{}).Where("code = ?", group.Code).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("配置组 Code 已存在")
	}

	return s.db.Create(group).Error
}

// GetConfigGroup 获取配置分组
func (s *ConfigServiceImpl) GetConfigGroup(id string) (*models.ConfigGroup, error) {
	if id == "" {
		return nil, errors.New("id 不能为空")
	}

	var group models.ConfigGroup
	err := s.db.Where("id = ?", id).First(&group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("配置组未找到")
		}
		return nil, err
	}

	return &group, nil
}

// UpdateConfigGroup 更新配置分组
func (s *ConfigServiceImpl) UpdateConfigGroup(group *models.ConfigGroup) error {
	if group == nil {
		return errors.New("配置组不能为 nil")
	}
	if group.ID == "" {
		return errors.New("id 不能为空")
	}

	// 检查分组是否存在
	var existingGroup models.ConfigGroup
	if err := s.db.Where("id = ?", group.ID).First(&existingGroup).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("配置组未找到")
		}
		return err
	}

	// 如果修改了code，检查新code是否与其他分组冲突
	if group.Code != existingGroup.Code {
		var count int64
		if err := s.db.Model(&models.ConfigGroup{}).Where("code = ? AND id != ?", group.Code, group.ID).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("配置组 Code 已存在")
		}
	}

	// 系统分组不允许修改某些字段
	if existingGroup.IsSystem {
		// 这里可以根据需要添加对系统分组字段的保护逻辑
	}

	return s.db.Save(group).Error
}

// DeleteConfigGroup 删除配置分组
func (s *ConfigServiceImpl) DeleteConfigGroup(id string) error {
	if id == "" {
		return errors.New("id 不能为空")
	}

	// 检查分组是否存在
	var group models.ConfigGroup
	if err := s.db.Where("id = ?", id).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("配置组未找到")
		}
		return err
	}

	// 系统分组不允许删除
	if group.IsSystem {
		return errors.New("系统配置组无法删除")
	}

	// 检查分组下是否有配置项
	var itemCount int64
	if err := s.db.Model(&models.ConfigItem{}).Where("group_id = ?", id).Count(&itemCount).Error; err != nil {
		return err
	}
	if itemCount > 0 {
		return errors.New("无法删除包含现有配置项的配置组")
	}

	return s.db.Where("id = ?", id).Delete(&models.ConfigGroup{}).Error
}

// ConfigGroupList 获取配置分组列表
func (s *ConfigServiceImpl) ConfigGroupList(req *config.ConfigGroupQueryRequest) (*query.PaginationResponse[models.ConfigGroup], error) {
	if req == nil {
		req = &config.ConfigGroupQueryRequest{}
		req.Normalize()
	}

	// 构建查询
	db := s.db.Model(&models.ConfigGroup{})

	// 应用查询条件
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		db = db.Where("name LIKE ? OR code LIKE ? OR description LIKE ?", keyword, keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 应用排序
	if sort := req.GetSort(); sort != "" {
		db = db.Order(sort)
	} else {
		// 默认排序
		db = db.Order("sort ASC, create_time DESC")
	}

	// 应用分页
	offset := req.GetOffset()
	var records []models.ConfigGroup
	if err := db.Offset(offset).Limit(req.Size).Find(&records).Error; err != nil {
		return nil, err
	}

	// 构建响应
	return query.BuildPaginationResponse(req, records, total), nil
}

// ============================================================
// ConfigItem CRUD 实现
// ============================================================

// CreateConfigItem 创建配置项
func (s *ConfigServiceImpl) CreateConfigItem(item *models.ConfigItem) error {
	if item == nil {
		return errors.New("配置项不能为 nil")
	}

	// 检查分组是否存在
	var group models.ConfigGroup
	if err := s.db.Where("id = ?", item.GroupID).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("配置组未找到")
		}
		return err
	}

	// 检查配置项代码是否已存在
	var count int64
	if err := s.db.Model(&models.ConfigItem{}).Where("code = ?", item.Code).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("配置项 Code 已存在")
	}

	return s.db.Create(item).Error
}

// GetConfigItem 获取配置项
func (s *ConfigServiceImpl) GetConfigItem(id string) (*models.ConfigItem, error) {
	if id == "" {
		return nil, errors.New("id 不能为空")
	}

	var item models.ConfigItem
	err := s.db.Where("id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("配置项未找到")
		}
		return nil, err
	}

	return &item, nil
}

// UpdateConfigItem 更新配置项
func (s *ConfigServiceImpl) UpdateConfigItem(item *models.ConfigItem) error {
	if item == nil {
		return errors.New("配置项不能为 nil")
	}
	if item.ID == "" {
		return errors.New("配置项 ID 不能为空")
	}

	// 检查配置项是否存在
	var existingItem models.ConfigItem
	if err := s.db.Where("id = ?", item.ID).First(&existingItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("配置项未找到")
		}
		return err
	}

	// 如果修改了分组ID，检查新分组是否存在
	if item.GroupID != existingItem.GroupID {
		var group models.ConfigGroup
		if err := s.db.Where("id = ?", item.GroupID).First(&group).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("配置组未找到")
			}
			return err
		}
	}

	// 如果修改了code，检查新code是否与其他配置项冲突
	if item.Code != existingItem.Code {
		var count int64
		if err := s.db.Model(&models.ConfigItem{}).Where("code = ? AND id != ?", item.Code, item.ID).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("配置项代码已存在")
		}
	}

	return s.db.Save(item).Error
}

// DeleteConfigItem 删除配置项
func (s *ConfigServiceImpl) DeleteConfigItem(id string) error {
	if id == "" {
		return errors.New("id 不能为空")
	}

	// 检查配置项是否存在
	var item models.ConfigItem
	if err := s.db.Where("id = ?", id).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("配置项未找到")
		}
		return err
	}

	return s.db.Where("id = ?", id).Delete(&models.ConfigItem{}).Error
}

// ConfigItemList 获取配置项列表
func (s *ConfigServiceImpl) ConfigItemList(req *config.ConfigItemQueryRequest) (*query.PaginationResponse[models.ConfigItem], error) {
	if req == nil {
		req = &config.ConfigItemQueryRequest{}
		req.Normalize()
	}

	// 构建查询
	db := s.db.Model(&models.ConfigItem{})

	// 应用查询条件
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		db = db.Where("name LIKE ? OR code LIKE ? OR description LIKE ?", keyword, keyword, keyword)
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	// 应用排序
	if sort := req.GetSort(); sort != "" {
		db = db.Order(sort)
	} else {
		// 默认排序
		db = db.Order("sort ASC, create_time DESC")
	}

	// 应用分页
	offset := req.GetOffset()
	var records []models.ConfigItem
	if err := db.Offset(offset).Limit(req.Size).Find(&records).Error; err != nil {
		return nil, err
	}

	// 构建响应
	return query.BuildPaginationResponse(req, records, total), nil
}
