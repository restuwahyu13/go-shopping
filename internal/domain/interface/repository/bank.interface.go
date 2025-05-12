package rinf

import "github.com/uptrace/bun"

type IBankRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
}
