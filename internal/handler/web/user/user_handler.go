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

	userPublicAssociatedInfo, err := h.userService.GetUserByID(userID)
	if err != nil {
		h.NotFound(c, "用户不存在")
		return
	}

	// 只返回公开信息
	h.Success(c, userPublicAssociatedInfo)
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	h.StartTimer(c)

	accountId, exists := c.Get("account_id")
	if !exists {
		h.Error(c, http.StatusUnauthorized, "认证失败")
	}

	associatedInfo, err := h.userService.GetUserProfile(accountId.(string))
	if err != nil {
		h.Error(c, http.StatusInternalServerError, "获取用户信息失败")
		return
	}

	h.Success(c, associatedInfo)
}
