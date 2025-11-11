package models

import (
	"galaxy/pkg/model"

	"gorm.io/datatypes"
)

// ============================================================
// 题目核心表
// ============================================================

// ProblemInfo 题目表
type ProblemInfo struct {
	model.BaseModel
	DisplayID    *string        `gorm:"column:display_id;type:varchar(32)"`
	CategoryID   string         `gorm:"column:category_id;type:varchar(32);default:0;index:idx_category_id"`
	Title        *string        `gorm:"column:title;type:varchar(255)"`
	Source       *string        `gorm:"column:source;type:varchar(255)"`
	URL          *string        `gorm:"column:url;type:varchar(255)"`
	TimeLimit    int            `gorm:"column:time_limit;default:0"`
	MemoryLimit  int            `gorm:"column:memory_limit;default:0"`
	Description  *string        `gorm:"column:description;type:text"`
	Languages    datatypes.JSON `gorm:"column:languages;type:jsonb"`
	Difficulty   int            `gorm:"column:difficulty;default:1;index:idx_difficulty"`
	Threshold    float64        `gorm:"column:threshold;type:decimal(10,2);default:0.50"`
	UseTemplate  bool           `gorm:"column:use_template;default:false"`
	CodeTemplate datatypes.JSON `gorm:"column:code_template;type:jsonb"`
	IsPublic     bool           `gorm:"column:is_public;default:false;index:idx_is_public"`
	IsVisible    bool           `gorm:"column:is_visible;default:true;index:idx_is_visible"`
	UseAI        bool           `gorm:"column:use_ai;default:false"`
}

func (ProblemInfo) TableName() string {
	return "problem_info"
}

// ProblemTagRel 题目标签关联表
type ProblemTagRel struct {
	model.BaseModel
	ProblemID string `gorm:"column:problem_id;type:varchar(32);not null"`
	TagID     string `gorm:"column:tag_id;type:varchar(32);not null"`
}

func (ProblemTagRel) TableName() string {
	return "problem_tag_rel"
}

// ProblemTestCase 题目测试用例表
type ProblemTestCase struct {
	model.BaseModel
	ProblemID      string  `gorm:"column:problem_id;type:varchar(32);not null;index:idx_problem_id"`
	CaseSign       *string `gorm:"column:case_sign;type:varchar(255)"`
	InputData      *string `gorm:"column:input_data;type:text"`
	ExpectedOutput *string `gorm:"column:expected_output;type:text"`
	InputFilePath  *string `gorm:"column:input_file_path;type:varchar(500)"`
	InputFileSize  int64   `gorm:"column:input_file_size;default:0"`
	OutputFilePath *string `gorm:"column:output_file_path;type:varchar(500)"`
	OutputFileSize int64   `gorm:"column:output_file_size;default:0"`
	IsSample       bool    `gorm:"column:is_sample;default:false"`
	Score          float64 `gorm:"column:score;type:decimal(10,2);default:0.00"`
}

func (ProblemTestCase) TableName() string {
	return "problem_test_case"
}
