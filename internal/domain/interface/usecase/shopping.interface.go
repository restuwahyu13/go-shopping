package uinf

import (
	"context"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type IShoppingUsecase interface {
	Checkout(ctx context.Context, body sdto.CheckoutDTO) hopt.Response
	CheckoutList(ctx context.Context) hopt.Response
}
