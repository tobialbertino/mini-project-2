package domain

type Pagiantion struct {
	Page       int
	PerPage    int
	Total      int
	TotalPages int
	Offset     int
}