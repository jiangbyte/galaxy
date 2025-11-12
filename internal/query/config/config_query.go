package config

import "galaxy/pkg/query"

// ConfigGroupQueryRequest 用户查询请求
type ConfigGroupQueryRequest struct {
	query.PaginationRequest
}

// ConfigItemQueryRequest 用户查询请求
type ConfigItemQueryRequest struct {
	query.PaginationRequest
}
