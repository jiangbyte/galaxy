package user

import (
	"galaxy/internal/service/web/user"
	"galaxy/pkg/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	handler.BaseHandler
	userService user.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: user.NewUserService(),
	}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	h.StartTimer(c)
	userID := c.Param("id")

	userInfo, err := h.userService.GetUserByID(userID)
	if err != nil {
		h.NotFound(c, "用户不存在")
		return
	}

	// 只返回公开信息
	h.Success(c, gin.H{
		"user": gin.H{
			"id":       userInfo.ID,
			"nickname": userInfo.Nickname,
			"avatar":   userInfo.Avatar,
			"title":    userInfo.Title,
			"level":    userInfo.Level,
		},
	})
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	h.StartTimer(c)

	accountId, exists := c.Get("account_id")
	if !exists {
		h.Error(c, http.StatusUnauthorized, "认证失败")
	}

	profile, err := h.userService.GetUserProfile(accountId.(string))
	if err != nil {
		h.Error(c, http.StatusInternalServerError, "获取用户信息失败")
		return
	}

	h.Success(c, profile)
}
