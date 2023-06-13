package tools

import (
	"strconv"
)

type Page[T any] struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"` //总数
	Pages       int `json:"pages"` //总页数
	Result      []T `json:"result"`
}

// 参数res：某个查询的所有数据库结果
// 参数currentPageString：url获取的参数，当前页
// 参数pageSizeString：url获取的参数，每页大小
func Pages[T any](res []T, currentPageString, pageSizeString string) Page[T] {
	currentPage, _ := strconv.Atoi(currentPageString)
	pageSize, _ := strconv.Atoi(pageSizeString)
	//分页后的结果集
	var result []T
	//截取起始位置offset
	offset := (currentPage - 1) * pageSize
	limit := pageSize
	//截取结束位置
	end := limit + offset
	//起始位大于数据库条数，返回空
	if offset >= len(res) {
		return Page[T]{}
	}
	//结束位置大于条数，截取结果到最后一条为止
	if end > len(res) {
		result = res[offset:]
	} else { //否则,取正常分页后的结果
		result = res[offset:end]
	}
	//没有结果集，原因：传入的res一条都没有
	if len(result) == 0 {
		return Page[T]{}
	}
	//取完结果计算响应的参数Total：总条数；Pages：总页数；result，分页后的结果集
	page := Page[T]{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Total:       len(res),
		Pages: func() int {
			//计算总页数
			pages := len(res) / pageSize
			if len(res)%pageSize == 0 {
				return pages
			} else {
				return pages + 1
			}
		}(),
		Result: result,
	}
	return page
}
