package pagination

type PaginationDb struct {
	Limit  int
	Offset int
}

type Pagination struct {
	Limit      int
	PageNumber int
}

type Meta struct {
	Pages        int
	PageNumber   int
	PageElements int
	All          int
}
