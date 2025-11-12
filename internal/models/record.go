package models

import (
	"galaxy/pkg/model"
	"time"

	"gorm.io/datatypes"
)

// ============================================================
// 用户学习记录表
// ============================================================

// RecordCodeLibrary 用户提交代码库
type RecordCodeLibrary struct {
	model.BaseModel
	UserID         *string        `gorm:"column:user_id;type:varchar(32)"`
	ModuleType     *string        `gorm:"column:module_type;type:varchar(32)"`
	ModuleID       *string        `gorm:"column:module_id;type:varchar(32)"`
	ProblemID      *string        `gorm:"column:problem_id;type:varchar(32);index:idx_problem_id"`
	SubmitID       *string        `gorm:"column:submit_id;type:varchar(32)"`
	SubmitTime     *time.Time     `gorm:"column:submit_time"`
	Language       *string        `gorm:"column:language;type:varchar(64);index:idx_language"`
	Code           *string        `gorm:"column:code;type:text"`
	CodeToken      datatypes.JSON `gorm:"column:code_token;type:jsonb"`
	CodeTokenName  datatypes.JSON `gorm:"column:code_token_name;type:jsonb"`
	CodeTokenTexts datatypes.JSON `gorm:"column:code_token_texts;type:jsonb"`
	CodeLength     int            `gorm:"column:code_length;default:0"`
	AccessCount    int            `gorm:"column:access_count;default:0"`
}

func (RecordCodeLibrary) TableName() string {
	return "record_code_library"
}

// RecordSolved 用户解决表
type RecordSolved struct {
	model.BaseModel
	ModuleType      *string    `gorm:"column:module_type;type:varchar(32)"`
	ModuleID        *string    `gorm:"column:module_id;type:varchar(32)"`
	UserID          *string    `gorm:"column:user_id;type:varchar(32)"`
	ProblemID       *string    `gorm:"column:problem_id;type:varchar(32)"`
	SubmitID        *string    `gorm:"column:submit_id;type:varchar(32)"`
	IsSolved        bool       `gorm:"column:is_solved;default:false"`
	FirstSolvedTime *time.Time `gorm:"column:first_solved_time"`
	SolvedTime      *time.Time `gorm:"column:solved_time"`
	FirstSubmitTime *time.Time `gorm:"column:first_submit_time"`
}

func (RecordSolved) TableName() string {
	return "record_solved"
}
