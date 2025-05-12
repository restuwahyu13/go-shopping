package uinf

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type IShoppingUsecase interface {
	CreateCheckoutCartShopping(ctx context.Context, req hdto.Request[[]sdto.CheckoutShoppingDTO]) hopt.Response
	ListCheckoutCartShopping(ctx context.Context) hopt.Response
}
