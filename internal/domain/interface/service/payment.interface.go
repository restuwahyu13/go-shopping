package sinf

import (
	"context"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type IPaymentService interface {
	PaymentCallbackSimulator(ctx context.Context, body any) hopt.Response
	PaymentWebhookSimulator(ctx context.Context, body any) hopt.Response
	PaymentSimulator(ctx context.Context, body any) hopt.Response
	PaymentStatus(ctx context.Context, params sdto.PaymentStatusDTO) hopt.Response
}
