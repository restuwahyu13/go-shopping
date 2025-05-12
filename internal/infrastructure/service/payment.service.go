package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null/v6/zero"
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

func (s PaymentService) PaymentCallbackSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment callback success"

	return res
}

func (s PaymentService) PaymentWebhookSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment webhook success"

	return res
}

func (s PaymentService) PaymentSimulator(ctx context.Context, req hdto.Request[any]) hopt.Response {
	res := hopt.Response{}
	res.StatCode = http.StatusOK
	res.Message = "Payment success"

	return res
}

func (s PaymentService) GeneratePayment(ctx context.Context, req hdto.Request[sdto.GeneratePaymentDTO]) hopt.Response {
	userId := fmt.Sprintf("%v", ctx.Value("user_id"))
	res := hopt.Response{}

	bankModel := new(model.BankModel)
	bankRepository := repo.NewBankRepository(ctx, s.DB)

	err := bankRepository.FindOne().Column("id").Where("LOWER(code) = ?", req.Body.Bank).Scan(ctx, bankModel)
	if err != nil && err != sql.ErrNoRows {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res

	} else if err == sql.ErrNoRows {
		res.StatCode = http.StatusNotFound
		res.ErrMsg = "Bank code is not exist in our system"

		return res
	}

	orderModel := new(model.OrderModel)
	orderRepository := repo.NewOrderRepository(ctx, s.DB)

	err = orderRepository.FindOne().Column("total_amount").Where("deleted_at IS NULL AND paid = false AND user_id = ?", userId).Order("created_at DESC").Scan(ctx, orderModel)
	if err != nil && err != sql.ErrNoRows {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res

	} else if err == sql.ErrNoRows || orderModel.TotalAmount != req.Body.Amount {
		res.StatCode = http.StatusUnprocessableEntity
		res.ErrMsg = "Miss match order amount"

		return res
	}

	paymentModel := new(model.PaymentModel)
	paymentRepository := repo.NewPaymentRepository(ctx, s.DB)

	err = paymentRepository.FindOne().Column("id").Where("deleted_at IS NULL AND verified_at IS NULL AND user_id = ? AND status = ?", userId, cons.CREATED).Scan(ctx, bankModel)
	if err != nil && err != sql.ErrNoRows {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res

	} else if err == sql.ErrNoRows {
		paymentRequesId := uuid.NewString()
		paymentInvoiceNumber := helper.NewRandom().Numeric(12)
		paymentExpiredAt := zero.TimeFrom(time.Now().Add(time.Duration(time.Second * 900)))

		paymentModel.UserID = userId
		paymentModel.BankID = bankModel.ID
		paymentModel.RequestID = paymentRequesId
		paymentModel.AccountNumber = paymentInvoiceNumber
		paymentModel.Amount = req.Body.Amount
		paymentModel.Method = req.Body.Method
		paymentModel.Status = cons.CREATED
		paymentModel.ExpiredAt = paymentExpiredAt

		paymentRepository := repo.NewPaymentRepository(ctx, s.DB)
		resutGeneratePayment, err := paymentRepository.Create(paymentModel).
			Returning("*",
				&paymentModel.ID,
				&paymentModel.RequestID,
				&paymentModel.AccountNumber,
				&paymentModel.Method,
				&paymentModel.Amount,
				&paymentModel.Status,
				&paymentModel.ExpiredAt,
				&paymentModel.CreatedAt,
			).Exec(ctx)

		if err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res

		} else if rows, err := resutGeneratePayment.RowsAffected(); rows < 1 || err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			if err == nil {
				res.StatCode = http.StatusPreconditionFailed
				res.ErrMsg = "Failed to make a generate payment"
			}

			return res
		}

		res.StatCode = http.StatusCreated
		res.Message = "Generate payment success"
		res.Data = paymentModel

		return res
	}

	res.StatCode = http.StatusPreconditionFailed
	res.ErrMsg = "You have already a payment, cancel to make a new payment"

	return res
}

func (s PaymentService) CheckStatusPayment(ctx context.Context, req hdto.Request[sdto.CheckStatusPaymentDTO]) hopt.Response {
	res := hopt.Response{}
	paymentModel := new(model.PaymentModel)

	paymentRepository := repo.NewPaymentRepository(ctx, s.DB)
	err := paymentRepository.FindOne().Column("id", "status").Where("deleted_at IS NULL AND id = ?", req.Param.ID).Scan(ctx, paymentModel)
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
