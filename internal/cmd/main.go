package main

import (
	"context"
	"restuwahyu13/shopping-cart/internal/cmd/api"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	cfg "restuwahyu13/shopping-cart/internal/infrastructure/common/config"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"

	"github.com/go-chi/chi/v5"
)

var (
	err error
	env *cdto.Environtment
)

func init() {
	env, err = cfg.NewEnvirontment(".env", ".", "env")
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}
}

func main() {
	ctx := context.Background()
	router := chi.NewRouter()

	db, err := cfg.Database(env)
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}

	rds, err := pkg.NewRedis(ctx, env.REDIS.URL)
	if err != nil {
		pkg.Logrus("fatal", err)
		return
	}

	app := api.NewApi(api.Api{ENV: env, ROUTER: router, DB: db, RDS: rds})
	app.Middleware()
	app.Router()
	app.Listener()
}
