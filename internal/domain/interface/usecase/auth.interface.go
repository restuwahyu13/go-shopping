package uinf

import (
	"context"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

type (
	IAuthUsecase interface {
		Login(ctx context.Context, body *sdto.LoginDTO) hopt.Response
	}
)
