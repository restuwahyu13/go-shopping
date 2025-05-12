package sinf

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type IPaymentService interface {
	CallbackSimulatorPayment(ctx context.Context, req hdto.Request[sdto.CallbackSimulatorPaymentDTO]) hopt.Response
	SimulatorPayment(ctx context.Context, req hdto.Request[sdto.SimulatorPaymentDTO]) hopt.Response
	GeneratePayment(ctx context.Context, req hdto.Request[sdto.GeneratePaymentDTO]) hopt.Response
	CheckStatusPayment(ctx context.Context, req hdto.Request[sdto.CheckStatusPaymentDTO]) hopt.Response
}
