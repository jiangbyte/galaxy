package user

import (
	"errors"
	"galaxy/internal/dto/user"
	"galaxy/internal/models"
	"galaxy/pkg/database"
	"gorm.io/gorm"
)

// UserService 用户服务接口定义
type UserService interface {
	GetUserByID(accountId string) (*user.UserPublicAssociatedInfo, error)
	GetUserByUsername(username string) (*models.AuthAccount, *user.UserPublicAssociatedInfo, error)
	GetUserProfile(accountId string) (*user.UserAssociatedProfile, error)

	// CreateUser 用户管理
	CreateUser(account *models.AuthAccount, userInfo *models.UserInfo) error
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
func (s *UserServiceImpl) CreateUser(account *models.AuthAccount, userInfo *models.UserInfo) error {
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
		preference := &models.UserPreference{
			AccountID: account.ID,
		}
		if err := tx.Create(preference).Error; err != nil {
			return err
		}

		// 创建用户档案
		profile := &models.UserProfile{
			AccountID: account.ID,
		}
		if err := tx.Create(profile).Error; err != nil {
			return err
		}

		// 创建用户统计
		stats := &models.UserStats{
			AccountID: account.ID,
		}
		return tx.Create(stats).Error
	})
}

// GetUserByID 根据ID获取用户信息
func (s *UserServiceImpl) GetUserByID(accountId string) (*user.UserPublicAssociatedInfo, error) {
	var userInfo user.UserPublicAssociatedInfo
	err := s.db.Where("account_id = ? AND deleted = false", accountId).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// GetUserByUsername 根据用户名获取用户
func (s *UserServiceImpl) GetUserByUsername(username string) (*models.AuthAccount, *user.UserPublicAssociatedInfo, error) {
	var account models.AuthAccount
	err := s.db.Where("username = ? AND deleted = false", username).First(&account).Error
	if err != nil {
		return nil, nil, err
	}

	// 使用预加载和选择特定字段
	var userPublicInfo user.UserPublicAssociatedInfo

	// 查询 UserInfo 相关字段
	var userInfo models.UserInfo
	err = s.db.Where("account_id = ? AND deleted = false", account.ID).First(&userInfo).Error
	if err != nil {
		return nil, nil, err
	}

	// 查询 UserProfile 相关字段
	var userProfile models.UserProfile
	err = s.db.Where("account_id = ? AND deleted = false", account.ID).First(&userProfile).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	// 查询 UserStats 相关字段
	var userStats models.UserStats
	err = s.db.Where("account_id = ? AND deleted = false", account.ID).First(&userStats).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	// 组装数据
	userPublicInfo = user.UserPublicAssociatedInfo{
		// UserInfo 字段
		AccountID:  userInfo.AccountID,
		Nickname:   userInfo.Nickname,
		Avatar:     userInfo.Avatar,
		Gender:     userInfo.Gender,
		Birthday:   userInfo.Birthday,
		Signature:  userInfo.Signature,
		Background: userInfo.Background,
		Interests:  userInfo.Interests,
		Website:    userInfo.Website,
		GitHub:     userInfo.GitHub,
		GitTee:     userInfo.GitTee,
		Blog:       userInfo.Blog,
		// UserProfile 字段
		Country:      userProfile.Country,
		Province:     userProfile.Province,
		City:         userProfile.City,
		ShowBirthday: userProfile.ShowBirthday,
		ShowLocation: userProfile.ShowLocation,
		// UserStats 字段
		Level:        userStats.Level,
		Exp:          userStats.Exp,
		TotalExp:     userStats.TotalExp,
		PostCount:    userStats.PostCount,
		CommentCount: userStats.CommentCount,
		LikeCount:    userStats.LikeCount,
		FollowCount:  userStats.FollowCount,
		FansCount:    userStats.FansCount,
	}

	return &account, &userPublicInfo, nil
}

// GetUserProfile 获取用户档案
func (s *UserServiceImpl) GetUserProfile(accountId string) (*user.UserAssociatedProfile, error) {
	var profile user.UserAssociatedProfile
	err := s.db.Where("account_id = ? AND deleted = false", accountId).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
