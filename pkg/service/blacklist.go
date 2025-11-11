// pkg/service/blacklist.go
package service

import (
	"fmt"
	"galaxy/pkg/redis"
	"galaxy/pkg/utils"
	"time"
)

type BlacklistService struct{}

var Blacklist = &BlacklistService{}

// AddToBlacklist 将token加入黑名单
func (s *BlacklistService) AddToBlacklist(tokenString string) error {
	// 解析token获取过期时间
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		return err
	}

	// 计算剩余过期时间
	expireTime := claims.ExpiresAt.Time
	now := time.Now()
	ttl := expireTime.Sub(now)

	if ttl <= 0 {
		// token已经过期，不需要加入黑名单
		return nil
	}

	// 使用token本身作为key，value可以是任意值，设置过期时间
	key := s.getBlacklistKey(tokenString)
	return redis.Set(key, "1", ttl)
}

// IsInBlacklist 检查token是否在黑名单中
func (s *BlacklistService) IsInBlacklist(tokenString string) (bool, error) {
	key := s.getBlacklistKey(tokenString)
	return redis.Exists(key)
}

// RemoveFromBlacklist 从黑名单移除token（通常不需要，因为会自动过期）
func (s *BlacklistService) RemoveFromBlacklist(tokenString string) error {
	key := s.getBlacklistKey(tokenString)
	return redis.Delete(key)
}

// 获取黑名单的Redis key
func (s *BlacklistService) getBlacklistKey(tokenString string) string {
	return fmt.Sprintf("jwt:blacklist:%s", tokenString)
}

// InvalidateUserTokens 使用户的所有token失效
func (s *BlacklistService) InvalidateUserTokens(userID string) error {
	// 存储失效时间戳
	key := s.getUserInvalidationKey(userID)
	invalidationTime := time.Now().Unix()

	// 设置为永久有效，或者设置一个很长的过期时间
	return redis.Set(key, invalidationTime, 365*24*time.Hour)
}

// GetUserInvalidationTime 获取用户token失效时间
func (s *BlacklistService) GetUserInvalidationTime(userID string) (int64, error) {
	key := s.getUserInvalidationKey(userID)
	result, err := redis.Get(key)
	if err != nil {
		// 检查是否是"key不存在"的错误
		if err.Error() == "redis: nil" {
			return 0, nil // 没有失效记录
		}
		return 0, err
	}

	// 将字符串转换为int64
	var invalidationTime int64
	_, err = fmt.Sscan(result, &invalidationTime)
	if err != nil {
		return 0, err
	}

	return invalidationTime, nil
}

func (s *BlacklistService) getUserInvalidationKey(userID string) string {
	return fmt.Sprintf("jwt:invalidation:%s", userID)
}
