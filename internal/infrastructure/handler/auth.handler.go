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

type AuthHandler struct {
	USECASE uinf.IAuthUsecase
}

func NewAuthHandler(options AuthHandler) hdinf.IAuthHandler {
	return AuthHandler{USECASE: options.USECASE}
}

func (h AuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var (
		req    sdto.LoginDTO   = sdto.LoginDTO{}
		res    hopt.Response   = hopt.Response{}
		ctx    context.Context = r.Context()
		parser hinf.IParser    = helper.NewParser()
	)

	if err := parser.Decode(r.Body, &req); err != nil {
		pkg.Logrus(cons.ERROR, err)
		helper.Api(rw, res)
		return
	}

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

	if res = h.USECASE.Login(ctx, &req); res.StatCode >= http.StatusBadRequest {
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
