package auth

import (
	"errors"
	"galaxy/internal/dto/user"
	"galaxy/internal/models"
	user2 "galaxy/internal/service/web/user"
	"galaxy/pkg/database"
	"galaxy/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// AuthService 用户服务接口定义
type AuthService interface {
	DoRegister(req *user.RegisterRequest) (*models.UserInfo, error)
	DoLogin(req *user.LoginRequest, clientIP string) (string, *user.UserPublicAssociatedInfo, error)
}

type AuthServiceImpl struct {
	db          *gorm.DB
	userService user2.UserService // 注入 UserService
}

func NewAuthService(userService user2.UserService) AuthService {
	return &AuthServiceImpl{
		db:          database.GetDB(),
		userService: userService,
	}
}

// DoRegister 用户注册
func (s *AuthServiceImpl) DoRegister(req *user.RegisterRequest) (*models.UserInfo, error) {
	// 检查用户名是否已存在
	var existingAccount models.AuthAccount
	err := s.db.Where("username = ? AND deleted = false", req.Username).First(&existingAccount).Error
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建用户账户
	account := &models.AuthAccount{
		Username: req.Username,
		Password: hashedPassword,
		Email:    &req.Email,
	}

	// 创建用户信息
	userInfo := &models.UserInfo{
		Nickname: req.Nickname,
	}

	if err := s.userService.CreateUser(account, userInfo); err != nil {
		return nil, errors.New("用户创建失败")
	}

	return userInfo, nil
}

// DoLogin 用户登录
func (s *AuthServiceImpl) DoLogin(req *user.LoginRequest, clientIP string) (string, *user.UserPublicAssociatedInfo, error) {
	// 获取用户
	account, userInfo, err := s.userService.GetUserByUsername(req.Username)
	if err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, account.Password) {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 生成 token
	token, err := utils.GenerateToken(userInfo.AccountID)
	if err != nil {
		return "", nil, errors.New("Token生成失败")
	}

	// 更新登录信息
	if err := s.updateLoginInfo(account.ID, clientIP); err != nil {
		return "", nil, errors.New("更新登录信息失败")
	}

	return token, userInfo, nil
}

func (s *AuthServiceImpl) updateLoginInfo(accountID, clientIP string) error {
	now := time.Now()

	updateData := models.LoginInfoUpdate{
		LastLoginTime: &now,
		LastLoginIP:   &clientIP,
		LockUntil:     nil,
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// 更新基础字段
		if err := tx.Table("auth_account").
			Where("id = ?", accountID).
			Updates(updateData).Error; err != nil {
			return err
		}

		// 更新计数字段
		return tx.Table("auth_account").
			Where("id = ?", accountID).
			Update("login_count", gorm.Expr("login_count + 1")).Error
	})
}
