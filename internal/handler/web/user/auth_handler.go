package user

import (
	user2 "galaxy/internal/dto/user"
	"galaxy/internal/service/share/auth"
	"galaxy/internal/service/web/user"
	"galaxy/pkg/captcha"
	"galaxy/pkg/handler"
	"galaxy/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	handler.BaseHandler
	authService auth.AuthService
}

func NewAuthHandler() *AuthHandler {
	// 创建 UserService 实例
	userService := user.NewUserService()
	// 注入到 AuthService
	authService := auth.NewAuthService(userService)
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) GenerateCaptcha(c *gin.Context) {
	h.StartTimer(c)

	// 生成验证码
	id, b64s, _, err := captcha.GenerateCaptcha()
	if err != nil {
		h.InternalServerError(c, "生成验证码失败")
		return
	}

	h.Success(c, user2.CaptchaResponse{
		CaptchaID:  id,
		CaptchaImg: b64s,
	})
}

func (h *AuthHandler) DoRegister(c *gin.Context) {
	h.StartTimer(c)

	var req user2.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	// 验证验证码
	if !captcha.VerifyCaptcha(req.CaptchaID, req.CaptchaCode) {
		h.BadRequest(c, "验证码错误")
		return
	}

	userInfo, err := h.authService.DoRegister(&req)
	if err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	h.Success(c, user2.RegisterResponse{
		UserID: userInfo.ID,
	})
}

func (h *AuthHandler) DoLogin(c *gin.Context) {
	h.StartTimer(c)

	var req user2.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	// 验证验证码
	if !captcha.VerifyCaptcha(req.CaptchaID, req.CaptchaCode) {
		h.BadRequest(c, "验证码错误")
		return
	}

	token, info, err := h.authService.DoLogin(&req, c.ClientIP())
	if err != nil {
		h.Unauthorized(c, err.Error())
		return
	}

	h.Success(c, user2.LoginResponse{
		Token: token,
		User:  *info,
	})
}

func (h *AuthHandler) DoLogout(c *gin.Context) {
	h.StartTimer(c)

	var req user2.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	// 将 token 加入黑名单
	if err := service.Blacklist.AddToBlacklist(req.Token); err != nil {
		h.InternalServerError(c, "注销失败")
		return
	}

	h.Success(c, user2.LogoutResponse{
		Message: "注销成功",
	})
}
