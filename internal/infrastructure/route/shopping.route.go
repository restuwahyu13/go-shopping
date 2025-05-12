package route

import (
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/handler"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/middleware"

	"github.com/go-chi/chi/v5"
)

type ShoppingRoute struct {
	ROUTER  chi.Router
	HANDLER hinf.IShoppingHandler
	ENV     *cdto.Environtment
	REDIS   pinf.IRedis
}

func NewShoppingRoute(options ShoppingRoute) {
	handler := ShoppingRoute{ROUTER: options.ROUTER, HANDLER: options.HANDLER}

	handler.ROUTER.Route(helper.Version("shopping"), func(r chi.Router) {
		r.Use(middleware.Auth(options.ENV.JWT.EXPIRED, options.REDIS))

		r.Post("/checkout", handler.HANDLER.CreateCheckoutCartShopping)
		r.Get("/checkout", handler.HANDLER.ListCheckoutCartShopping)
	})
}
