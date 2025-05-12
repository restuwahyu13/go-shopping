package usecase

import (
	"context"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	uinf "restuwahyu13/shopping-cart/internal/domain/interface/usecase"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type AuthUsecase struct {
	SERVICE sinf.IAuthService
}

func NewAuthUsecase(options AuthUsecase) uinf.IAuthUsecase {
	return AuthUsecase{SERVICE: options.SERVICE}
}

func (u AuthUsecase) Login(ctx context.Context, body *sdto.LoginDTO) hopt.Response {
	return u.SERVICE.Login(ctx, body)
}
