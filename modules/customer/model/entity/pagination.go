package entity

type Pagination struct {
	Page       int
	PerPage    int
	Total      int
	TotalPages int
	Offset     int
}
