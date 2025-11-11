package models

import (
	"galaxy/pkg/model"
)

// ============================================================
// 内容管理表
// ============================================================

// CmsArticle 系统文章表
type CmsArticle struct {
	model.BaseModel
	Title      string  `gorm:"column:title;type:varchar(255)"`
	Subtitle   *string `gorm:"column:subtitle;type:varchar(255)"`
	Cover      *string `gorm:"column:cover;type:varchar(255)"`
	Author     *string `gorm:"column:author;type:varchar(255)"`
	Summary    *string `gorm:"column:summary;type:varchar(255)"`
	Sort       int     `gorm:"column:sort;type:smallint;default:99"`
	ToURL      *string `gorm:"column:to_url;type:varchar(255)"`
	ParentID   string  `gorm:"column:parent_id;type:varchar(32);default:0"`
	Type       string  `gorm:"column:type;type:varchar(32);default:0"`
	Category   string  `gorm:"column:category;type:varchar(32);default:0"`
	Content    *string `gorm:"column:content;type:text"`
	ModuleType string  `gorm:"column:module_type;type:varchar(255)"`
}

func (CmsArticle) TableName() string {
	return "cms_article"
}

// CmsBanner 横幅表
type CmsBanner struct {
	model.BaseModel
	Title             string  `gorm:"column:title;type:varchar(255)"`
	Banner            *string `gorm:"column:banner;type:varchar(255)"`
	ButtonText        *string `gorm:"column:button_text;type:varchar(255)"`
	IsVisibleButton   bool    `gorm:"column:is_visible_button;default:true"`
	JumpModule        *string `gorm:"column:jump_module;type:varchar(255)"`
	JumpType          *string `gorm:"column:jump_type;type:varchar(255)"`
	JumpTarget        *string `gorm:"column:jump_target;type:varchar(255)"`
	TargetBlank       bool    `gorm:"column:target_blank;default:false"`
	Sort              int     `gorm:"column:sort;type:smallint;default:99"`
	Subtitle          *string `gorm:"column:subtitle;type:varchar(255)"`
	IsVisibleSubtitle bool    `gorm:"column:is_visible_subtitle;default:true"`
	IsVisible         bool    `gorm:"column:is_visible;default:true"`
	ModuleType        string  `gorm:"column:module_type;type:varchar(255)"`
}

func (CmsBanner) TableName() string {
	return "cms_banner"
}

// CmsCategory 分类表
type CmsCategory struct {
	model.BaseModel
	Name       string `gorm:"column:name;type:varchar(255)"`
	IsVisible  bool   `gorm:"column:is_visible;default:true"`
	ModuleType string `gorm:"column:module_type;type:varchar(255)"`
	ParentID   string `gorm:"column:parent_id;type:varchar(32);default:0"`
}

func (CmsCategory) TableName() string {
	return "cms_category"
}

// CmsNotice 公告表
type CmsNotice struct {
	model.BaseModel
	Title      string  `gorm:"column:title;type:varchar(64)"`
	Cover      *string `gorm:"column:cover;type:varchar(255)"`
	URL        *string `gorm:"column:url;type:varchar(255)"`
	Sort       int     `gorm:"column:sort;type:smallint;default:99"`
	Content    *string `gorm:"column:content;type:text"`
	IsVisible  bool    `gorm:"column:is_visible;default:true"`
	ModuleType string  `gorm:"column:module_type;type:varchar(255)"`
}

func (CmsNotice) TableName() string {
	return "cms_notice"
}

// CmsTag 标签表
type CmsTag struct {
	model.BaseModel
	Name       string `gorm:"column:name;type:varchar(255)"`
	IsVisible  bool   `gorm:"column:is_visible;default:true"`
	ModuleType string `gorm:"column:module_type;type:varchar(255)"`
	ParentID   string `gorm:"column:parent_id;type:varchar(32);default:0"`
}

func (CmsTag) TableName() string {
	return "cms_tag"
}
