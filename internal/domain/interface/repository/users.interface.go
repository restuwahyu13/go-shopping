package rinf

import "github.com/uptrace/bun"

type IUserRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
}
