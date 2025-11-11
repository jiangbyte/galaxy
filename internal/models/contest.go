package models

import (
	"galaxy/pkg/model"
	"time"

	"gorm.io/datatypes"
)

// ============================================================
// 竞赛与活动表
// ============================================================

// ContestInfo 竞赛表
type ContestInfo struct {
	model.BaseModel
	Title             string         `gorm:"column:title;type:varchar(255);not null"`
	Description       *string        `gorm:"column:description;type:text"`
	ContestType       *string        `gorm:"column:contest_type;type:varchar(50)"`
	RuleType          *string        `gorm:"column:rule_type;type:varchar(50)"`
	Category          *string        `gorm:"column:category;type:varchar(32)"`
	Cover             *string        `gorm:"column:cover;type:varchar(255)"`
	MaxTeamMembers    int            `gorm:"column:max_team_members;default:1"`
	IsTeamContest     bool           `gorm:"column:is_team_contest;default:false"`
	IsVisible         bool           `gorm:"column:is_visible;default:true"`
	IsPublic          bool           `gorm:"column:is_public;default:false"`
	Password          *string        `gorm:"column:password;type:varchar(100)"`
	RegisterStartTime *time.Time     `gorm:"column:register_start_time"`
	RegisterEndTime   *time.Time     `gorm:"column:register_end_time"`
	ContestStartTime  *time.Time     `gorm:"column:contest_start_time"`
	ContestEndTime    *time.Time     `gorm:"column:contest_end_time"`
	FrozenTime        int            `gorm:"column:frozen_time;default:0"`
	PenaltyTime       int            `gorm:"column:penalty_time;default:20"`
	AllowedLanguages  datatypes.JSON `gorm:"column:allowed_languages;type:jsonb"`
	Status            *string        `gorm:"column:status;type:varchar(32)"`
	Sort              int            `gorm:"column:sort;default:0"`
}

func (ContestInfo) TableName() string {
	return "contest_info"
}

// ContestAuth 竞赛认证表
type ContestAuth struct {
	model.BaseModel
	ContestID string `gorm:"column:contest_id;type:varchar(32);not null;index:idx_contest_id"`
	UserID    string `gorm:"column:user_id;type:varchar(32);not null;index:idx_user_id"`
	IsAuth    bool   `gorm:"column:is_auth;default:false"`
}

func (ContestAuth) TableName() string {
	return "contest_auth"
}

// ContestParticipant 竞赛参与表
type ContestParticipant struct {
	model.BaseModel
	ContestID    string     `gorm:"column:contest_id;type:varchar(32);not null;index:idx_contest_id"`
	UserID       string     `gorm:"column:user_id;type:varchar(32);not null;index:idx_user_id"`
	TeamID       *string    `gorm:"column:team_id;type:varchar(32)"`
	TeamName     *string    `gorm:"column:team_name;type:varchar(255)"`
	IsTeamLeader bool       `gorm:"column:is_team_leader;default:false"`
	RegisterTime *time.Time `gorm:"column:register_time"`
	Status       *string    `gorm:"column:status;type:varchar(32)"`
}

func (ContestParticipant) TableName() string {
	return "contest_participant"
}

// ContestProblem 竞赛题目表
type ContestProblem struct {
	model.BaseModel
	ContestID   string  `gorm:"column:contest_id;type:varchar(32);not null;index:idx_contest_id"`
	DisplayID   *string `gorm:"column:display_id;type:varchar(32)"`
	ProblemCode string  `gorm:"column:problem_code;type:varchar(10);not null;index:idx_problem_code"`
	ProblemID   string  `gorm:"column:problem_id;type:varchar(32);not null"`
	Score       int     `gorm:"column:score;default:0"`
	Sort        int     `gorm:"column:sort;default:0"`
}

func (ContestProblem) TableName() string {
	return "contest_problem"
}
