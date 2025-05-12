package rinf

import (
	ropt "restuwahyu13/shopping-cart/internal/domain/output/repository"

	"github.com/uptrace/bun"
)

type IProductConfigRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
	FindProductConfigByProductId(productId string) (*ropt.FindProductConfigByProductId, error)
}
