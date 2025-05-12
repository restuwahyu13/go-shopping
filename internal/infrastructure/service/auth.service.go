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
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"

	"github.com/uptrace/bun"
)

type AuthService struct {
	ENV *cdto.Environtment
	DB  *bun.DB
	RDS pinf.IRedis
}

func NewAuthService(options AuthService) sinf.IAuthService {
	return AuthService{
		DB:  options.DB,
		ENV: options.ENV,
		RDS: options.RDS,
	}
}

func (s AuthService) Login(ctx context.Context, body *sdto.LoginDTO) hopt.Response {
	res := hopt.Response{}
	userModel := new(model.UsersModel)

	userRepository := repo.NewUsersRepository(s.DB)
	err := userRepository.FindOne().Column("id", "verified_at").Where("email = ?", body.Email).Scan(ctx, userModel)

	if err != nil && err != sql.ErrNoRows {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res
	}

	if err == sql.ErrNoRows {
		res.StatCode = http.StatusNotFound
		res.ErrMsg = "User account is not registered"

		return res
	} else if !userModel.VerifiedAt.Valid {
		res.StatCode = http.StatusNotFound
		res.ErrMsg = "User account is not verified"

		return res
	}

	jwt := pkg.NewJsonWebToken(s.ENV, s.RDS)
	resultSignature, err := jwt.Sign(userModel.ID, userModel)

	if err != nil {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res
	}

	res.StatCode = http.StatusOK
	res.Message = "Login success"
	res.Data = resultSignature

	return res
}
