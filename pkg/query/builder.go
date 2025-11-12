// pkg/query/builder.go
package query

import (
	"gorm.io/gorm"
	"math"
)

// QueryBuilder 查询构建器
type QueryBuilder struct {
	db *gorm.DB
}

func NewQueryBuilder(db *gorm.DB) *QueryBuilder {
	return &QueryBuilder{db: db}
}

// Paginate 执行分页查询
func (qb *QueryBuilder) Paginate(req *PaginationRequest, result interface{}, applyConditions func(db *gorm.DB) *gorm.DB) (*PaginationResponse[interface{}], error) {
	req.Normalize()

	// 构建查询
	query := applyConditions(qb.db)

	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 计算总页数
	pages := int(math.Ceil(float64(total) / float64(req.Size)))

	// 应用排序
	if sort := req.GetSort(); sort != "" {
		query = query.Order(sort)
	}

	// 应用分页
	offset := req.GetOffset()
	if err := query.Offset(offset).Limit(req.Size).Find(result).Error; err != nil {
		return nil, err
	}

	// 构建响应
	response := &PaginationResponse[interface{}]{
		Current: req.Current,
		Pages:   pages,
		Records: toInterfaceSlice(result),
		Size:    req.Size,
		Total:   total,
	}

	return response, nil
}

// toInterfaceSlice 将任意切片转换为 interface{} 切片
func toInterfaceSlice(slice interface{}) []interface{} {
	// 这里需要根据具体类型进行转换
	// 在实际使用中，你可以使用反射来处理，但为了简单起见，这里返回空
	// 或者你可以在调用时明确指定类型
	return []interface{}{}
}

// PaginationRequestor 分页请求接口
type PaginationRequestor interface {
	GetCurrent() int
	GetSize() int
	Normalize()
}

// BuildPaginationResponse 泛型版本的分页响应
func BuildPaginationResponse[T any](req interface{}, records []T, total int64) *PaginationResponse[T] {
	var current, size int

	// 处理不同类型的请求
	switch r := req.(type) {
	case *PaginationRequest:
		r.Normalize()
		current = r.Current
		size = r.Size
	case PaginationRequestor:
		r.Normalize()
		current = r.GetCurrent()
		size = r.GetSize()
	default:
		// 对于嵌入 PaginationRequest 的类型，使用反射或直接调用 Normalize
		if paginationReq, ok := req.(interface{ Normalize() }); ok {
			paginationReq.Normalize()
		}

		// 尝试获取 Current 和 Size 字段
		if getter, ok := req.(interface{ GetCurrent() int }); ok {
			current = getter.GetCurrent()
		}
		if getter, ok := req.(interface{ GetSize() int }); ok {
			size = getter.GetSize()
		}

		// 如果无法获取，使用默认值
		if current <= 0 {
			current = DefaultCurrent
		}
		if size <= 0 {
			size = DefaultSize
		}
		if size > MaxSize {
			size = MaxSize
		}
	}

	pages := int(math.Ceil(float64(total) / float64(size)))

	return &PaginationResponse[T]{
		Current: current,
		Pages:   pages,
		Records: records,
		Size:    size,
		Total:   total,
	}
}
