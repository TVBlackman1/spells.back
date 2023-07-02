package pagination

import "math"

type PaginationDb struct {
	Limit  int
	Offset int
}

type Pagination struct {
	Limit      int
	PageNumber int
}

type Meta struct {
	Pages      int
	PageNumber int
	PageLimit  int
	All        int
}

func GetMeta(limit int, count int, pageNumber int) Meta {
	return Meta{
		Pages:      int(math.Ceil(float64(count) / float64(limit))),
		PageNumber: pageNumber,
		PageLimit:  limit,
		All:        count,
	}
}
