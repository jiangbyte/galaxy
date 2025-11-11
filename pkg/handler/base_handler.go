package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

const startTimeKey = "startTime"

// StartTimer 开始计时（在处理器开始时调用）
func (h *BaseHandler) StartTimer(c *gin.Context) {
	c.Set(startTimeKey, time.Now())
}

// GetElapsedTime 获取经过的时间（在处理器结束时调用）
func (h *BaseHandler) GetElapsedDuration(c *gin.Context) time.Duration {
	if startTime, exists := c.Get(startTimeKey); exists {
		if start, ok := startTime.(time.Time); ok {
			return time.Since(start)
		}
	}
	return 0
}

// GetElapsedTime 获取经过的时间（返回带单位的字符串）
func (h *BaseHandler) GetElapsedTime(c *gin.Context) string {
	if startTime, exists := c.Get(startTimeKey); exists {
		if start, ok := startTime.(time.Time); ok {
			elapsed := time.Since(start)
			return h.formatDuration(elapsed)
		}
	}
	return "0ms"
}

// formatDuration 格式化时间间隔为可读字符串
func (h *BaseHandler) formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return d.String() // 纳秒
	} else if d < time.Millisecond {
		return d.Round(time.Microsecond).String() // 微秒
	} else if d < time.Second {
		return d.Round(time.Millisecond).String() // 毫秒
	} else {
		return d.Round(time.Millisecond).String() // 秒（保留毫秒精度）
	}
}

func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": "success",
		"success": true,
		"time":    h.GetElapsedTime(c), // 返回毫秒
	})
}

func (h *BaseHandler) SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"data":    data,
		"message": message,
		"success": true,
		"time":    h.GetElapsedTime(c),
	})
}

func (h *BaseHandler) Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"data":    nil,
		"message": message,
		"success": false,
		"time":    h.GetElapsedTime(c),
	})
	c.Abort()
}

func (h *BaseHandler) BadRequest(c *gin.Context, message string) {
	h.Error(c, http.StatusBadRequest, message)
}

func (h *BaseHandler) Unauthorized(c *gin.Context, message string) {
	h.Error(c, http.StatusUnauthorized, message)
}

func (h *BaseHandler) InternalServerError(c *gin.Context, message string) {
	h.Error(c, http.StatusInternalServerError, message)
}

func (h *BaseHandler) NotFound(c *gin.Context, message string) {
	h.Error(c, http.StatusNotFound, message)
}
