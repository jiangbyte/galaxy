// pkg/query/pagination.go
package query

import (
	"strings"
)

// PaginationRequest 分页请求参数
type PaginationRequest struct {
	Current   int    `json:"current" form:"current"`       // 当前页码
	Size      int    `json:"size" form:"size"`             // 每页条数
	SortField string `json:"sort_field" form:"sort_field"` // 排序字段
	SortOrder string `json:"sort_order" form:"sort_order"` // 排序方式
	Keyword   string `json:"keyword" form:"keyword"`       // 关键词
}

// PaginationResponse 分页响应
type PaginationResponse[T any] struct {
	Current int   `json:"current"` // 当前页码
	Pages   int   `json:"pages"`   // 总页数
	Records []T   `json:"records"` // 数据记录
	Size    int   `json:"size"`    // 每页条数
	Total   int64 `json:"total"`   // 总记录数
}

// 默认值
const (
	DefaultCurrent = 1
	DefaultSize    = 10
	MaxSize        = 1000
)

// Normalize 规范化分页参数
func (p *PaginationRequest) Normalize() {
	if p.Current <= 0 {
		p.Current = DefaultCurrent
	}
	if p.Size <= 0 {
		p.Size = DefaultSize
	}
	if p.Size > MaxSize {
		p.Size = MaxSize
	}
	if p.SortOrder != "ascend" && p.SortOrder != "descend" {
		p.SortOrder = "descend" // 默认降序
	}
}

// GetOffset 获取偏移量
func (p *PaginationRequest) GetOffset() int {
	return (p.Current - 1) * p.Size
}

// GetSort 获取排序字符串
func (p *PaginationRequest) GetSort() string {
	if p.SortField == "" {
		return ""
	}

	order := "DESC"
	if p.SortOrder == "ascend" {
		order = "ASC"
	}

	// 简单的字段名安全过滤
	safeField := strings.ReplaceAll(p.SortField, ";", "")
	safeField = strings.ReplaceAll(safeField, "'", "")
	safeField = strings.ReplaceAll(safeField, "\"", "")

	return safeField + " " + order
}
