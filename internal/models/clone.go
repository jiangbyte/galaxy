package models

import (
	"galaxy/pkg/model"
	"time"

	"gorm.io/datatypes"
)

// ============================================================
// 代码相似度表
// ============================================================

// SimilarityStat 相似统计表
type SimilarityStat struct {
	model.BaseModel
	StatisticsType         int            `gorm:"column:statistics_type;default:0"`
	TaskID                 *string        `gorm:"column:task_id;type:varchar(32);index:idx_task_id"`
	ModuleType             *string        `gorm:"column:module_type;type:varchar(32)"`
	ModuleID               *string        `gorm:"column:module_id;type:varchar(32)"`
	ProblemID              *string        `gorm:"column:problem_id;type:varchar(32);index:idx_problem_id"`
	SampleCount            int            `gorm:"column:sample_count;default:0"`
	SimilarityGroupCount   int            `gorm:"column:similarity_group_count;default:0"`
	AvgSimilarity          float64        `gorm:"column:avg_similarity;type:decimal(10,2);default:0.00"`
	MaxSimilarity          float64        `gorm:"column:max_similarity;type:decimal(10,2);default:0.00"`
	Threshold              float64        `gorm:"column:threshold;type:decimal(10,2);default:0.50"`
	SimilarityDistribution datatypes.JSON `gorm:"column:similarity_distribution;type:jsonb"`
	DegreeStatistics       datatypes.JSON `gorm:"column:degree_statistics;type:jsonb"`
}

func (SimilarityStat) TableName() string {
	return "similarity_stat"
}

// SimilaritySegment 相似片段表
type SimilaritySegment struct {
	model.BaseModel
	TaskID     *string `gorm:"column:task_id;type:varchar(32)"`
	TaskType   bool    `gorm:"column:task_type;default:false"`
	ProblemID  *string `gorm:"column:problem_id;type:varchar(32)"`
	ModuleType *string `gorm:"column:module_type;type:varchar(32)"`
	ModuleID   *string `gorm:"column:module_id;type:varchar(32)"`
	Language   *string `gorm:"column:language;type:varchar(64)"`
	Similarity float64 `gorm:"column:similarity;type:decimal(10,2);default:0.00"`

	SubmitUser           *string        `gorm:"column:submit_user;type:varchar(32)"`
	SubmitCode           *string        `gorm:"column:submit_code;type:text"`
	SubmitCodeLength     int            `gorm:"column:submit_code_length;default:0"`
	SubmitID             *string        `gorm:"column:submit_id;type:varchar(32)"`
	SubmitTime           *time.Time     `gorm:"column:submit_time"`
	SubmitCodeToken      datatypes.JSON `gorm:"column:submit_code_token;type:jsonb"`
	SubmitCodeTokenName  datatypes.JSON `gorm:"column:submit_code_token_name;type:jsonb"`
	SubmitCodeTokenTexts datatypes.JSON `gorm:"column:submit_code_token_texts;type:jsonb"`

	LibraryUser           *string        `gorm:"column:library_user;type:varchar(32)"`
	LibraryCode           *string        `gorm:"column:library_code;type:text"`
	LibraryCodeLength     int            `gorm:"column:library_code_length;default:0"`
	LibraryID             *string        `gorm:"column:library_id;type:varchar(32)"`
	LibraryTime           *time.Time     `gorm:"column:library_time"`
	LibraryCodeToken      datatypes.JSON `gorm:"column:library_code_token;type:jsonb"`
	LibraryCodeTokenName  datatypes.JSON `gorm:"column:library_code_token_name;type:jsonb"`
	LibraryCodeTokenTexts datatypes.JSON `gorm:"column:library_code_token_texts;type:jsonb"`
}

func (SimilaritySegment) TableName() string {
	return "similarity_segment"
}
