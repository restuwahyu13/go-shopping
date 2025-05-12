package helper

import (
	"github.com/uptrace/bun"

	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
)

type sql struct{}

func NewSql() hinf.ISql {
	return &sql{}
}

func (h sql) IncColumn(cols ...string) func(sq *bun.SelectQuery) *bun.SelectQuery {
	return func(sq *bun.SelectQuery) *bun.SelectQuery {
		sq.Column(cols...)
		return sq
	}
}

func (h sql) ExcColumn(cols ...string) func(sq *bun.SelectQuery) *bun.SelectQuery {
	return func(sq *bun.SelectQuery) *bun.SelectQuery {
		sq.ExcludeColumn(cols...)
		return sq
	}
}

func (h sql) Column(options *hdto.ColumnOptions) func(sq *bun.SelectQuery) *bun.SelectQuery {
	return func(sq *bun.SelectQuery) *bun.SelectQuery {

		if len(options.Inc) > 0 {
			sq.Column(options.Inc...)
		}

		if len(options.Exc) > 0 {
			sq.ExcludeColumn(options.Exc...)
		}

		return sq
	}
}
