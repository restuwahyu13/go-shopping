package usecase

import (
	"context"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	uinf "restuwahyu13/shopping-cart/internal/domain/interface/usecase"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type PaymentUsecase struct {
	SERVICE sinf.IPaymentService
}

func NewPaymentUsecase(options PaymentUsecase) uinf.IPaymentUsecase {
	return PaymentUsecase{SERVICE: options.SERVICE}
}

func (u PaymentUsecase) PaymentCallbackSimulator(ctx context.Context, body any) hopt.Response {
	return u.SERVICE.PaymentCallbackSimulator(ctx, body)
}

func (u PaymentUsecase) PaymentWebhookSimulator(ctx context.Context, body any) hopt.Response {
	return u.SERVICE.PaymentWebhookSimulator(ctx, body)
}

func (u PaymentUsecase) PaymentSimulator(ctx context.Context, body any) hopt.Response {
	return u.SERVICE.PaymentSimulator(ctx, body)
}

func (u PaymentUsecase) PaymentStatus(ctx context.Context, params sdto.PaymentStatusDTO) hopt.Response {
	return u.SERVICE.PaymentStatus(ctx, params)
}
