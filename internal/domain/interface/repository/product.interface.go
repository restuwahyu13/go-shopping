package rinf

import "github.com/uptrace/bun"

type IProductRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
}
