package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogger 自定义 Gin 日志中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 读取请求体（用于日志记录）
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 处理请求
		c.Next()

		// 结束时间
		timestamp := time.Now()
		latency := timestamp.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		bodyStr := string(bodyBytes)
		if len(bodyStr) > 500 { // 限制日志长度
			bodyStr = bodyStr[:500] + "..."
		}

		// 根据状态码选择日志级别
		event := log.Info()
		if statusCode >= 400 && statusCode < 500 {
			event = log.Warn()
		} else if statusCode >= 500 {
			event = log.Error()
		}

		event.
			Str("🌐", "HTTP").
			Int("status", statusCode).
			Str("method", method).
			Str("path", path).
			Str("query", raw).
			Str("ip", clientIP).
			Str("latency", latency.String()).
			//Str("user-agent", c.Request.UserAgent()).
			Str("time", timestamp.Format("2006-01-02 15:04:05")).
			Msg(fmt.Sprintf("%s %s", method, path))

		// 记录请求体
		if os.Getenv("GIN_MODE") == "debug" {
			if len(bodyStr) > 0 && bodyStr != "{}" {
				log.Debug().
					Str("🌐", "HTTP_BODY").
					Str("method", method).
					Str("path", path).
					Str("body", bodyStr).
					Msg("Request body")
			}
		}

		if errorMessage != "" {
			log.Error().
				Str("🌐", "HTTP_ERROR").
				Str("method", method).
				Str("path", path).
				Str("error", errorMessage).
				Msg("Request error")
		}
	}
}

// GinRecovery 自定义恢复中间件
func GinRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(io.Discard, func(c *gin.Context, err interface{}) {
		log.Error().
			Str("🌐", "HTTP_PANIC").
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Interface("error", err).
			Msg("HTTP panic recovered")

		c.AbortWithStatus(500)
	})
}
