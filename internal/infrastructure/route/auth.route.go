package route

import (
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/handler"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"

	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	ROUTER  chi.Router
	HANDLER hinf.IAuthHandler
}

func NewAuthRoute(options AuthRoute) {
	handler := AuthRoute{ROUTER: options.ROUTER, HANDLER: options.HANDLER}

	handler.ROUTER.Route(helper.Version("auth"), func(r chi.Router) {
		r.Post("/login", handler.HANDLER.Login)
	})
}
