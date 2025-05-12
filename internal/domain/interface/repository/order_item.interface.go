package rinf

import "github.com/uptrace/bun"

type IOrderItemRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
}
