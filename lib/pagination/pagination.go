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
	Pages         int `json:"pages,omitempty"`
	PageNumber    int `json:"page_number,omitempty"`
	PageLimit     int `json:"page_limit,omitempty"`
	OnCurrentPage int `json:"on_current_page,omitempty"`
	AllRecords    int `json:"all_records,omitempty"`
}

func GetMeta(limit int, count int, pageNumber int) Meta {
	pages := int(math.Ceil(float64(count) / float64(limit)))
	recordsOnCurrentPage := 0
	if pages == pageNumber {
		// from 0 to limit; if 0 -- its not last page or onCurrentPage=limit
		recordsOnCurrentPage = count % limit
	}
	if recordsOnCurrentPage == 0 {
		recordsOnCurrentPage = limit
	}
	return Meta{
		Pages:         int(math.Ceil(float64(count) / float64(limit))),
		PageNumber:    pageNumber,
		PageLimit:     limit,
		OnCurrentPage: recordsOnCurrentPage,
		AllRecords:    count,
	}
}
