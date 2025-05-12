package hopt

type Pagination struct {
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
	Perpage float64 `json:"perpage"`
	Total   int     `json:"total"`
}
