package pagination

import (
	"database/sql"
	"math"
)

// 分页请求数据
type Paging struct {
	Page  int `json:"page"`  // 页码
	Limit int `json:"limit"` // 每页条数
	Total int `json:"total"` // 总数据条数
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (p *Paging) TotalPage() int {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := p.Total / p.Limit
	if p.Total%p.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}

type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

// 分页返回数据
type PageResult struct {
	Page    *Paging     `json:"page"`    // 分页信息
	Results interface{} `json:"results"` // 数据
}

// Cursor分页返回数据
type CursorResult struct {
	Results interface{} `json:"results"` // 数据
	Cursor  string      `json:"cursor"`  // 下一页
}

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

type Pagination struct {
	Total       uint        `json:"total"`
	PerPage     uint        `json:"per_page"`
	CurrentPage uint        `json:"current_page"`
	LastPage    uint        `json:"last_page"`
	From        uint        `json:"from"`
	To          uint        `json:"to"`
	Data        interface{} `json:"data"`
}

func Paginate(page, pageSize, count uint, data interface{}) Pagination {
	var lastPage uint = 1

	to := page * pageSize

	if to > count {
		to = count
	}

	from := (page-1)*pageSize + 1

	if count == 0 || from > count {
		return Pagination{PerPage: pageSize, CurrentPage: page, LastPage: lastPage, Data: data}
	}

	lastPage = uint(math.Ceil(float64(count) / float64(pageSize)))

	return Pagination{
		Total: count, PerPage: pageSize, CurrentPage: page,
		LastPage: lastPage, From: from, To: to, Data: data,
	}
}

func CanPaginate(page, pageSize, count uint) bool {
	if page == 0 {
		page = 1
	}

	from := (page-1)*pageSize + 1
	if count == 0 || from > count {
		return false
	}

	return true
}

func PurePageArgs(page uint, pageSize uint) (uint, uint) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	return page, pageSize
}

func EmptyPagination(page, pageSize uint) Pagination {
	return Paginate(page, pageSize, 0, make([]interface{}, 0))
}
