package domain

import "math"

type Pagination struct {
	Pages  uint64 `json:"pages"`
	Total  uint64 `json:"total"`
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offest"`
}

func NewPagination(limit, offset, total uint64) Pagination {
	pagination := Pagination{
		Pages:  0,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
	if limit <= 0 {
		return pagination
	}

	pagination.Pages = uint64(math.Ceil(float64(total) / float64(limit)))

	return pagination
}
