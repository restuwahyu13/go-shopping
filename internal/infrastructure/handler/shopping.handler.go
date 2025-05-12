package handler

import (
	"context"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hdinf "restuwahyu13/shopping-cart/internal/domain/interface/handler"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	uinf "restuwahyu13/shopping-cart/internal/domain/interface/usecase"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"

	gpc "github.com/restuwahyu13/go-playground-converter"
)

type ShoppingHandler struct {
	USECASE uinf.IShoppingUsecase
}

func NewShoppingHandler(options ShoppingHandler) hdinf.IShoppingHandler {
	return ShoppingHandler{USECASE: options.USECASE}
}

func (h ShoppingHandler) Checkout(rw http.ResponseWriter, r *http.Request) {
	var (
		body   sdto.CheckoutDTO = sdto.CheckoutDTO{}
		res    hopt.Response    = hopt.Response{}
		ctx    context.Context  = r.Context()
		parser hinf.IParser     = helper.NewParser()
	)

	if err := parser.Decode(r.Body, &body); err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

	validator, err := gpc.Validator(body)
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

	if res = h.USECASE.Checkout(ctx, body); res.StatCode >= http.StatusBadRequest {
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

func (h ShoppingHandler) CheckoutList(rw http.ResponseWriter, r *http.Request) {
	var (
		res hopt.Response   = hopt.Response{}
		ctx context.Context = r.Context()
	)

	if res = h.USECASE.CheckoutList(ctx); res.StatCode >= http.StatusBadRequest {
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
