package handler

import (
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
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
	ctx := r.Context()

	res := hopt.Response{}
	req := hdto.Request[any]{}

	if res = h.USECASE.PaymentCallbackSimulator(ctx, req); res.StatCode >= http.StatusBadRequest {
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
	ctx := r.Context()

	res := hopt.Response{}
	req := hdto.Request[any]{}

	if res = h.USECASE.PaymentWebhookSimulator(ctx, req); res.StatCode >= http.StatusBadRequest {
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

func (h PaymentHandler) SimulatorPayment(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	parser := helper.NewParser()

	res := hopt.Response{}
	req := hdto.Request[sdto.SimulatorPaymentDTO]{}

	if err := parser.Decode(r.Body, &req.Body); err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

	validator, err := gpc.Validator(req.Body)
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

	if res = h.USECASE.SimulatorPayment(ctx, req); res.StatCode >= http.StatusBadRequest {
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

func (h PaymentHandler) GeneratePayment(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	parser := helper.NewParser()

	res := hopt.Response{}
	req := hdto.Request[sdto.GeneratePaymentDTO]{}

	if err := parser.Decode(r.Body, &req.Body); err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

	validator, err := gpc.Validator(req.Body)
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

	if res = h.USECASE.GeneratePayment(ctx, req); res.StatCode >= http.StatusBadRequest {
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

func (h PaymentHandler) CheckStatusPayment(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res := hopt.Response{}
	req := hdto.Request[sdto.CheckStatusPaymentDTO]{}

	req.Param.ID = chi.URLParam(r, "id")
	validator, err := gpc.Validator(req.Param)
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

	if res = h.USECASE.CheckStatusPayment(ctx, req); res.StatCode >= http.StatusBadRequest {
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
