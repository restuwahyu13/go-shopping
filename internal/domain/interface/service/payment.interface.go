package sinf

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type IPaymentService interface {
	PaymentCallbackSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response
	PaymentWebhookSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response
	PaymentSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response
	GeneratePayment(ctx context.Context, req hdto.Request[sdto.GeneratePaymentDTO]) hopt.Response
	CheckStatusPayment(ctx context.Context, req hdto.Request[sdto.CheckStatusPaymentDTO]) hopt.Response
}
