package usecase

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
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

func (u PaymentUsecase) CallbackSimulatorPayment(ctx context.Context, req hdto.Request[sdto.CallbackSimulatorPaymentDTO]) hopt.Response {
	return u.SERVICE.CallbackSimulatorPayment(ctx, req)
}

func (u PaymentUsecase) SimulatorPayment(ctx context.Context, req hdto.Request[sdto.SimulatorPaymentDTO]) hopt.Response {
	return u.SERVICE.SimulatorPayment(ctx, req)
}

func (u PaymentUsecase) GeneratePayment(ctx context.Context, req hdto.Request[sdto.GeneratePaymentDTO]) hopt.Response {
	return u.SERVICE.GeneratePayment(ctx, req)
}

func (u PaymentUsecase) CheckStatusPayment(ctx context.Context, req hdto.Request[sdto.CheckStatusPaymentDTO]) hopt.Response {
	return u.SERVICE.CheckStatusPayment(ctx, req)
}
