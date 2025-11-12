package user

// ====================== 注册认证请求与响应 ======================

// RegisterRequest 注册请求
type RegisterRequest struct {
	Nickname string `json:"nickname" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`

	CaptchaID   string `json:"captcha_id" binding:"required"`   // 验证码ID
	CaptchaCode string `json:"captcha_code" binding:"required"` // 验证码
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

// ====================== 登录认证请求与响应 ======================

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

	CaptchaID   string `json:"captcha_id" binding:"required"`   // 验证码ID
	CaptchaCode string `json:"captcha_code" binding:"required"` // 验证码
}

// LoginResponse 定义登录响应结构体
type LoginResponse struct {
	Token string                   `json:"token"`
	User  UserPublicAssociatedInfo `json:"user"`
}

// ====================== 验证码响应 ======================

type CaptchaResponse struct {
	CaptchaID  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
}

// ====================== 注销 ======================

// LogoutRequest 注销请求
type LogoutRequest struct {
	Token string `json:"token" binding:"required"`
}

// LogoutResponse 注销响应
type LogoutResponse struct {
	Message string `json:"message"`
}
