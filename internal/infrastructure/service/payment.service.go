package service

import (
	"context"
	"database/sql"
	"net/http"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"

	"github.com/uptrace/bun"
)

type PaymentService struct {
	ENV *cdto.Environtment
	DB  *bun.DB
	RDS pinf.IRedis
}

func NewPaymentService(options PaymentService) sinf.IPaymentService {
	return PaymentService{
		ENV: options.ENV,
		DB:  options.DB,
		RDS: options.RDS,
	}
}

func (s PaymentService) PaymentCallbackSimulator(ctx context.Context, body any) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment callback success"

	return res
}

func (s PaymentService) PaymentWebhookSimulator(ctx context.Context, body any) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment webhook success"

	return res
}

func (s PaymentService) PaymentSimulator(ctx context.Context, body any) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment success"

	return res
}

func (s PaymentService) PaymentStatus(ctx context.Context, params sdto.PaymentStatusDTO) hopt.Response {
	res := hopt.Response{}

	paymentModel := new(model.PaymentModel)
	paymentRepository := repo.NewPaymentRepository(ctx, s.DB)

	err := paymentRepository.FindOne().Column("id", "status").Where("deleted_at IS NULL AND id = ?", params.ID).Scan(ctx, paymentModel)
	if err != nil && err != sql.ErrNoRows {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res
	}

	if err == sql.ErrNoRows {
		res.StatCode = http.StatusNotFound
		res.ErrMsg = "Payment status is not exist in our system"

		return res
	}

	res.StatCode = http.StatusOK
	res.Message = "Success"
	res.Data = paymentModel

	return res
}
