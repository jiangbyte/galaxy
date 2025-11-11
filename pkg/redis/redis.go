package redis

import (
	"context"
	"fmt"
	"galaxy/pkg/config"
	"galaxy/pkg/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client
var ctx = context.Background()

func Init() {
	cfg := config.Get().Redis

	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		// 禁用高级特性
		DisableIndentity: true, // 禁用身份识别
		Protocol:         2,    // 使用 RESP2 协议
	})

	// 测试连接
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Redis().
			Str("host", cfg.Host).
			Int("port", cfg.Port).
			Err(err).
			Msg("Connection failed")
		panic(err)
	}

	logger.Connected("Redis")
	logger.Redis().
		Str("host", cfg.Host).
		Int("port", cfg.Port).
		Int("db", cfg.DB).
		Msg("Connection details")
}

func GetClient() *redis.Client {
	if client == nil {
		logger.Error().
			Msg("Redis not initialized. Call Init() first.")
	}
	return client
}

// 队列操作 (单机部署时使用)
func Enqueue(queueName string, message []byte) error {
	return client.LPush(ctx, queueName, message).Err()
}

func Dequeue(queueName string) ([]byte, error) {
	result, err := client.BRPop(ctx, 0, queueName).Result()
	if err != nil {
		return nil, err
	}
	if len(result) < 2 {
		return nil, fmt.Errorf("invalid queue result")
	}
	return []byte(result[1]), nil
}

// 缓存操作
func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func Delete(key string) error {
	return client.Del(ctx, key).Err()
}

func Exists(key string) (bool, error) {
	result, err := client.Exists(ctx, key).Result()
	return result > 0, err
}
