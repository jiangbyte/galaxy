package user

import (
	"galaxy/internal/models"
	"gorm.io/datatypes"
	"time"
)

// UserPublicAssociatedInfo 用户公开关联信息
type UserPublicAssociatedInfo struct {
	// USER INFO
	AccountID  string         `json:"account_id"`
	Nickname   string         `json:"nickname"`
	Avatar     *string        `json:"avatar"`
	Gender     int            `json:"gender"`
	Birthday   *time.Time     `json:"birthday"`
	Signature  *string        `json:"signature"`
	Background *string        `json:"background"`
	Interests  datatypes.JSON `json:"interests"`
	Website    *string        `json:"website"`
	GitHub     *string        `json:"github"`
	GitTee     *string        `json:"gitee"`
	Blog       *string        `json:"blog"`
	// USER PROFILE
	Country      *string `json:"country"`
	Province     *string `json:"province"`
	City         *string `json:"city"`
	ShowBirthday bool    `json:"show_birthday"`
	ShowLocation bool    `json:"show_location"`
	// UserStats
	Level        int   `json:"level"`
	Exp          int64 `json:"exp"`
	TotalExp     int64 `json:"total_exp"`
	PostCount    int64 `json:"post_count"`
	CommentCount int64 `json:"comment_count"`
	LikeCount    int64 `json:"like_count"`
	FollowCount  int64 `json:"follow_count"`
	FansCount    int64 `json:"fans_count"`
}

// UserAssociatedProfile 用户公开关联信息
type UserAssociatedProfile struct {
	models.UserInfo
	models.UserProfile
}
