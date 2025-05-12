package usecase

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	uinf "restuwahyu13/shopping-cart/internal/domain/interface/usecase"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type ShoppingUsecase struct {
	SERVICE sinf.IShoppingService
}

func NewShoppingUsecase(options ShoppingUsecase) uinf.IShoppingUsecase {
	return ShoppingUsecase{SERVICE: options.SERVICE}
}

func (u ShoppingUsecase) CreateCheckoutCartShopping(ctx context.Context, req hdto.Request[[]sdto.CheckoutShoppingDTO]) hopt.Response {
	return u.SERVICE.CreateCheckoutCartShopping(ctx, req)
}

func (u ShoppingUsecase) ListCheckoutCartShopping(ctx context.Context) hopt.Response {
	return u.SERVICE.ListCheckoutCartShopping(ctx)
}
