package captcha

import (
	"galaxy/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"time"
)

var store base64Captcha.Store
var driver base64Captcha.Driver

func Init() {
	// 使用 Redis 存储验证码
	store = NewRedisStore()

	// 配置验证码驱动
	driver = base64Captcha.NewDriverDigit(
		80,  // 高度
		240, // 宽度
		5,   // 验证码长度
		0.2, // 最大倾斜角度
		40,  // 干扰线数量
	)
}

// RedisStore 实现 base64Captcha.Store 接口
type RedisStore struct {
	prefix string
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		prefix: "captcha:",
	}
}

func (s *RedisStore) Set(id string, value string) error {
	key := s.prefix + id
	return redis.Set(key, value, 5*time.Minute) // 5分钟过期
}

func (s *RedisStore) Get(id string, clear bool) string {
	key := s.prefix + id
	value, err := redis.Get(key)
	if err != nil {
		return ""
	}

	if clear {
		_ = redis.Delete(key)
	}
	return value
}

func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	storedValue := s.Get(id, clear)
	return storedValue == answer
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha() (string, string, string, error) {
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return "", "", "", err
	}

	return id, b64s, answer, nil
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id, code string) bool {
	return store.Verify(id, code, true) // 验证后清除
}
