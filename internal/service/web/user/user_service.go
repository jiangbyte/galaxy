package user

import (
	"galaxy/internal/models"
	"galaxy/pkg/database"
	"gorm.io/gorm"
)

// UserService 用户服务接口定义
type UserService interface {
	// GetUserByID 用户信息查询
	GetUserByID(accountId string) (*models.AuthUserInfo, error)
	GetUserByUsername(username string) (*models.AuthAccount, *models.AuthUserInfo, error)
	GetUserProfile(accountId string) (*models.AuthUserProfile, error)

	// CreateUser 用户管理
	CreateUser(account *models.AuthAccount, userInfo *models.AuthUserInfo) error
}

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	db *gorm.DB
}

// 确保 UserServiceImpl 实现 UserService 接口
var _ UserService = (*UserServiceImpl)(nil)

func NewUserService() UserService {
	return &UserServiceImpl{
		db: database.GetDB(),
	}
}

// CreateUser 创建用户
func (s *UserServiceImpl) CreateUser(account *models.AuthAccount, userInfo *models.AuthUserInfo) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 创建账户
		if err := tx.Create(account).Error; err != nil {
			return err
		}

		// 创建用户信息
		userInfo.AccountID = account.ID
		if err := tx.Create(userInfo).Error; err != nil {
			return err
		}

		// 创建用户偏好设置（默认值）
		preference := &models.AuthUserPreference{
			AccountID: account.ID,
		}
		if err := tx.Create(preference).Error; err != nil {
			return err
		}

		// 创建用户统计
		stats := &models.AuthUserStats{
			AccountID: account.ID,
		}
		return tx.Create(stats).Error
	})
}

// GetUserByID 根据ID获取用户信息
func (s *UserServiceImpl) GetUserByID(accountId string) (*models.AuthUserInfo, error) {
	var userInfo models.AuthUserInfo
	err := s.db.Where("account_id = ? AND deleted = false", accountId).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// GetUserByUsername 根据用户名获取用户
func (s *UserServiceImpl) GetUserByUsername(username string) (*models.AuthAccount, *models.AuthUserInfo, error) {
	var account models.AuthAccount
	err := s.db.Where("username = ? AND deleted = false", username).First(&account).Error
	if err != nil {
		return nil, nil, err
	}

	var userInfo models.AuthUserInfo
	err = s.db.Where("account_id = ? AND deleted = false", account.ID).First(&userInfo).Error
	if err != nil {
		return nil, nil, err
	}

	return &account, &userInfo, nil
}

// GetUserProfile 获取用户档案
func (s *UserServiceImpl) GetUserProfile(accountId string) (*models.AuthUserProfile, error) {
	var profile models.AuthUserProfile
	err := s.db.Where("account_id = ? AND deleted = false", accountId).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
