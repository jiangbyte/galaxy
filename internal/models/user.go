package models

import (
	"galaxy/pkg/model"
	"gorm.io/datatypes"
	"time"
)

// ============================================================
// 用户信息
// ============================================================

// UserInfo 用户基本信息表
type UserInfo struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	// 基础身份信息
	Nickname string     `gorm:"column:nickname;type:varchar(128);not null;index:idx_user_nickname"` // 昵称
	Avatar   *string    `gorm:"column:avatar;type:varchar(255)"`                                    // 头像
	Gender   int        `gorm:"column:gender;type:smallint;default:0"`                              // 性别 0: 未知 1: 男 2: 女
	Birthday *time.Time `gorm:"column:birthday;type:date"`                                          // 生日
	// 展示信息
	Signature  *string        `gorm:"column:signature;type:varchar(500)"`  // 个性签名
	Background *string        `gorm:"column:background;type:varchar(255)"` // 个人背景图片
	Interests  datatypes.JSON `gorm:"column:interests;type:jsonb"`         // 兴趣标签
	// 社交链接
	Website *string `gorm:"column:website;type:varchar(255)"` // 个人网站
	GitHub  *string `gorm:"column:github;type:varchar(100)"`  // GitHub
	GitTee  *string `gorm:"column:gitee;type:varchar(100)"`   // GitTee
	Blog    *string `gorm:"column:blog;type:varchar(255)"`    // 博客
}

func (UserInfo) TableName() string {
	return "user_info"
}

// UserProfile 用户档案详情表
type UserProfile struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	// 教育职业信息
	RealName  *string `gorm:"column:real_name;type:varchar(64)"`  // 真实姓名
	School    *string `gorm:"column:school;type:varchar(100)"`    // 学校
	Major     *string `gorm:"column:major;type:varchar(100)"`     //  专业
	StudentID *string `gorm:"column:student_id;type:varchar(50)"` // 学号
	Company   *string `gorm:"column:company;type:varchar(100)"`   //  公司
	JobTitle  *string `gorm:"column:job_title;type:varchar(100)"` // 职位
	Industry  *string `gorm:"column:industry;type:varchar(100)"`  // 行业
	// 地理位置
	Country  *string `gorm:"column:country;type:varchar(50)"`   // 国家
	Province *string `gorm:"column:province;type:varchar(50)"`  // 省份
	City     *string `gorm:"column:city;type:varchar(50)"`      //  城市
	Location *string `gorm:"column:location;type:varchar(100)"` // 详细地址
	// 社交信息
	QQ     *string `gorm:"column:qq;type:varchar(20)"`     // QQ
	WeChat *string `gorm:"column:wechat;type:varchar(50)"` // 微信
	// 隐私设置
	ShowBirthday bool `gorm:"column:show_birthday;default:false"` // 是否显示生日
	ShowLocation bool `gorm:"column:show_location;default:true"`  // 是否显示地理位置
}

func (UserProfile) TableName() string {
	return "user_profile"
}

// UserPreference 用户偏好设置表
type UserPreference struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	// 界面设置
	Theme    string `gorm:"column:theme;type:varchar(50);default:light"`    // 主题
	Language string `gorm:"column:language;type:varchar(10);default:zh-CN"` // 系统语言
	// 通知设置
	EmailNotifications bool `gorm:"column:email_notifications;default:true"` // 邮件通知
	PushNotifications  bool `gorm:"column:push_notifications;default:true"`  // 推送通知
	// 隐私与展示
	AllowDirectMessage bool `gorm:"column:allow_direct_message;default:true"` // 允许私信
}

func (UserPreference) TableName() string {
	return "user_preference"
}

// UserStats 用户统计信息表
type UserStats struct {
	model.BaseModel
	AccountID string `gorm:"column:account_id;type:varchar(32);not null;uniqueIndex:idx_user_info_account"`
	// 等级与经验
	Level    int   `gorm:"column:level;default:1;index:idx_user_level"` // 等级
	Exp      int64 `gorm:"column:exp;default:0"`                        // 经验值
	TotalExp int64 `gorm:"column:total_exp;default:0"`                  // 累计经验值
	// 活跃统计
	LoginDays           int `gorm:"column:login_days;default:0"`            // 登录天数
	ContinuousLoginDays int `gorm:"column:continuous_login_days;default:0"` // 连续登录天数
	// 内容统计
	PostCount    int64 `gorm:"column:post_count;default:0"`    // 发帖数
	CommentCount int64 `gorm:"column:comment_count;default:0"` // 评论数
	LikeCount    int64 `gorm:"column:like_count;default:0"`    // 获赞数
	FollowCount  int64 `gorm:"column:follow_count;default:0"`  // 关注数
	FansCount    int64 `gorm:"column:fans_count;default:0"`    // 粉丝数
}

func (UserStats) TableName() string {
	return "user_stats"
}
