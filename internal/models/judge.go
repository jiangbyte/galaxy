package models

import (
	"galaxy/pkg/model"
)

// ============================================================
// 提交与判题表
// ============================================================

// JudgeCase 判题结果用例表
type JudgeCase struct {
	model.BaseModel
	SubmitID       string  `gorm:"column:submit_id;type:varchar(32);not null"`
	CaseSign       *string `gorm:"column:case_sign;type:varchar(255)"`
	InputData      *string `gorm:"column:input_data;type:text"`
	OutputData     *string `gorm:"column:output_data;type:text"`
	ExpectedOutput *string `gorm:"column:expected_output;type:text"`
	InputFilePath  *string `gorm:"column:input_file_path;type:varchar(500)"`
	InputFileSize  int64   `gorm:"column:input_file_size;default:0"`
	OutputFilePath *string `gorm:"column:output_file_path;type:varchar(500)"`
	OutputFileSize int64   `gorm:"column:output_file_size;default:0"`
	MaxTime        float64 `gorm:"column:max_time;type:decimal(10,2);default:0.00"`
	MaxMemory      float64 `gorm:"column:max_memory;type:decimal(10,2);default:0.00"`
	IsSample       bool    `gorm:"column:is_sample;default:false"`
	Score          float64 `gorm:"column:score;type:decimal(10,2);default:0.00"`
	Status         *string `gorm:"column:status;type:varchar(32)"`
	Message        *string `gorm:"column:message;type:text"`
	ExitCode       int64   `gorm:"column:exit_code;default:0"`
}

func (JudgeCase) TableName() string {
	return "judge_case"
}

// JudgeSubmit 提交表
type JudgeSubmit struct {
	model.BaseModel
	UserID        *string `gorm:"column:user_id;type:varchar(32);index:idx_user_id"`
	ModuleType    *string `gorm:"column:module_type;type:varchar(32)"`
	ModuleID      *string `gorm:"column:module_id;type:varchar(32)"`
	ProblemID     *string `gorm:"column:problem_id;type:varchar(32);index:idx_problem_id"`
	Language      *string `gorm:"column:language;type:varchar(64);index:idx_language"`
	Code          *string `gorm:"column:code;type:text"`
	CodeLength    int     `gorm:"column:code_length;default:0"`
	IsTestSubmit  bool    `gorm:"column:is_test_submit;default:false"`
	IsAdminSubmit bool    `gorm:"column:is_admin_submit;default:false"`
	MaxTime       int     `gorm:"column:max_time;default:0"`
	MaxMemory     int     `gorm:"column:max_memory;default:0"`
	Message       *string `gorm:"column:message;type:text"`
	Status        *string `gorm:"column:status;type:varchar(32)"`
	IsFinish      bool    `gorm:"column:is_finish;default:false"`
	TaskID        *string `gorm:"column:task_id;type:varchar(32)"`
}

func (JudgeSubmit) TableName() string {
	return "judge_submit"
}
