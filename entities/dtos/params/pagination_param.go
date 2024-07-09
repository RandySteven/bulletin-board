package params

import (
	"net/url"
	"strconv"
)

type PaginationParam struct {
	Page int `json:"page"`
}

func GetPagination(queryParam url.Values) PaginationParam {
	page, _ := strconv.Atoi(queryParam.Get("page"))
	return PaginationParam{
		Page: page,
	}
}

func (pagination PaginationParam) GetLimitOffset() (limit int, offset int) {
	limit = 10
	offset = 0
	if pagination.Page == 0 || pagination.Page == 1 {
		return limit, offset
	}
	offset = (limit * pagination.Page) - limit
	return limit, offset
}
