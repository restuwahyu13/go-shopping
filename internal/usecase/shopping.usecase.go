package usecase

import (
	"context"
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

func (u ShoppingUsecase) Checkout(ctx context.Context, body sdto.CheckoutDTO) hopt.Response {
	return u.SERVICE.Checkout(ctx, body)
}

func (u ShoppingUsecase) CheckoutList(ctx context.Context) hopt.Response {
	return u.SERVICE.CheckoutList(ctx)
}
