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

	gpc "github.com/restuwahyu13/go-playground-converter"
)

type ShoppingHandler struct {
	USECASE uinf.IShoppingUsecase
}

func NewShoppingHandler(options ShoppingHandler) hdinf.IShoppingHandler {
	return ShoppingHandler{USECASE: options.USECASE}
}

func (h ShoppingHandler) CreateCheckoutCartShopping(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	parser := helper.NewParser()

	res := hopt.Response{}
	req := hdto.Request[[]sdto.CheckoutShoppingDTO]{}

	if err := parser.Decode(r.Body, &req.Body); err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

	for _, body := range req.Body {
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
	}

	if res = h.USECASE.CreateCheckoutCartShopping(ctx, req); res.StatCode >= http.StatusBadRequest {
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

func (h ShoppingHandler) ListCheckoutCartShopping(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := hopt.Response{}

	if res = h.USECASE.ListCheckoutCartShopping(ctx); res.StatCode >= http.StatusBadRequest {
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
