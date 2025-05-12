package helper

import (
	"math"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

func Pagination(limit, offset, total int) *hopt.Pagination {
	res := new(hopt.Pagination)

	res.Limit = limit
	res.Offset = offset
	res.Perpage = math.Ceil(float64(total) / float64(limit))
	res.Total = total

	return res
}
