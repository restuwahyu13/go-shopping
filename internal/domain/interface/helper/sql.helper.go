package hinf

import (
	"github.com/uptrace/bun"

	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
)

type ISql interface {
	IncColumn(cols ...string) func(sq *bun.SelectQuery) *bun.SelectQuery
	ExcColumn(cols ...string) func(sq *bun.SelectQuery) *bun.SelectQuery
	Column(options *hdto.ColumnOptions) func(sq *bun.SelectQuery) *bun.SelectQuery
}
