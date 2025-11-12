package system

import (
	"galaxy/pkg/database"
	"gorm.io/gorm"
)

type SystemService struct {
	db *gorm.DB
}

func NewSystemService() *SystemService {
	return &SystemService{
		db: database.GetDB(),
	}
}

//// GetAllSysConfigs 获取所有系统配置
//func (s *SystemService) GetAllSysConfigs() ([]models.SysConfig, error) {
//	var configs []models.SysConfig
//	err := s.db.Where("deleted = false").Order("sort ASC").Find(&configs).Error
//	return configs, err
//}
//
//// UpdateSysConfig 更新系统配置
//func (s *SystemService) UpdateSysConfig(code, value string) error {
//	return s.db.Model(&models.SysConfig{}).Where("code = ? AND deleted = false", code).Update("value", value).Error
//}
//
//// GetSysDict 获取系统字典
//func (s *SystemService) GetSysDict(dictType string) ([]models.SysDict, error) {
//	var dicts []models.SysDict
//	err := s.db.Where("dict_type = ? AND deleted = false", dictType).Order("sort_order ASC").Find(&dicts).Error
//	return dicts, err
//}
//
//// GetSysMenus 获取系统菜单
//func (s *SystemService) GetSysMenus() ([]models.SysMenu, error) {
//	var menus []models.SysMenu
//	err := s.db.Where("deleted = false AND visible = true").Order("sort ASC").Find(&menus).Error
//	return menus, err
//}
