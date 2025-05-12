package handler

import (
	"context"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hdinf "restuwahyu13/shopping-cart/internal/domain/interface/handler"
	uinf "restuwahyu13/shopping-cart/internal/domain/interface/usecase"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"

	"github.com/go-chi/chi/v5"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type PaymentHandler struct {
	USECASE uinf.IPaymentUsecase
}

func NewPaymentHandler(options PaymentHandler) hdinf.IPaymentHandler {
	return PaymentHandler{USECASE: options.USECASE}
}

func (h PaymentHandler) PaymentCallbackSimulator(rw http.ResponseWriter, r *http.Request) {
	var (
		res hopt.Response   = hopt.Response{}
		ctx context.Context = r.Context()
	)

	if res = h.USECASE.PaymentCallbackSimulator(ctx, nil); res.StatCode >= http.StatusBadRequest {
		if res.StatCode >= http.StatusInternalServerError {
			pkg.Logrus(cons.ERROR, res.ErrMsg)
			res.ErrMsg = cons.DEFAULT_ERR_MSG
		}

		helper.Api(rw, res)
		return
	}

	helper.Api(rw, res)
	return
}

func (h PaymentHandler) PaymentWebhookSimulator(rw http.ResponseWriter, r *http.Request) {
	var (
		res hopt.Response   = hopt.Response{}
		ctx context.Context = r.Context()
	)

	if res = h.USECASE.PaymentWebhookSimulator(ctx, nil); res.StatCode >= http.StatusBadRequest {
		if res.StatCode >= http.StatusInternalServerError {
			pkg.Logrus(cons.ERROR, res.ErrMsg)
			res.ErrMsg = cons.DEFAULT_ERR_MSG
		}

		helper.Api(rw, res)
		return
	}

	helper.Api(rw, res)
	return
}

func (h PaymentHandler) PaymentSimulator(rw http.ResponseWriter, r *http.Request) {
	var (
		res hopt.Response   = hopt.Response{}
		ctx context.Context = r.Context()
	)

	if res = h.USECASE.PaymentSimulator(ctx, nil); res.StatCode >= http.StatusBadRequest {
		if res.StatCode >= http.StatusInternalServerError {
			pkg.Logrus(cons.ERROR, res.ErrMsg)
			res.ErrMsg = cons.DEFAULT_ERR_MSG
		}

		helper.Api(rw, res)
		return
	}

	helper.Api(rw, res)
	return
}

func (h PaymentHandler) PaymentStatus(rw http.ResponseWriter, r *http.Request) {
	var (
		res hopt.Response         = hopt.Response{}
		req sdto.PaymentStatusDTO = sdto.PaymentStatusDTO{}
		ctx context.Context       = r.Context()
	)

	req.ID = chi.URLParam(r, "id")

	validator, err := gpc.Validator(req)
	if err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

	if validator != nil {
		res.StatCode = http.StatusUnprocessableEntity
		res.Errors = validator.Errors

		helper.Api(rw, res)
		return
	}

	if res = h.USECASE.PaymentStatus(ctx, req); res.StatCode >= http.StatusBadRequest {
		if res.StatCode >= http.StatusInternalServerError {
			pkg.Logrus(cons.ERROR, res.ErrMsg)
			res.ErrMsg = cons.DEFAULT_ERR_MSG
		}

		helper.Api(rw, res)
		return
	}

	helper.Api(rw, res)
	return
}
