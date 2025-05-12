package rinf

import (
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	ropt "restuwahyu13/shopping-cart/internal/domain/output/repository"

	"github.com/uptrace/bun"
)

type IProductItemRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create() *bun.InsertQuery
	Update() *bun.UpdateQuery
	FindCheckoutProductItemByPromotionRules(trx bun.Tx, promotionConfig *ropt.FindProductConfigByProductId, checkouts []sdto.CheckoutShoppingDTO) (*ropt.FindCheckoutProductItemPromotionRules, error)
}
