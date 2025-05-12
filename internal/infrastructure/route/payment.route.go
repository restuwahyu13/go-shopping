package route

import (
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/handler"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/middleware"

	"github.com/go-chi/chi/v5"
)

type PaymentRoute struct {
	ROUTER  chi.Router
	HANDLER hinf.IPaymentHandler
	ENV     *cdto.Environtment
	REDIS   pinf.IRedis
}

func NewPaymentRoute(options PaymentRoute) {
	handler := PaymentRoute{ROUTER: options.ROUTER, HANDLER: options.HANDLER}

	handler.ROUTER.Route(helper.Version("simulator"), func(r chi.Router) {
		r.Post("/callback", handler.HANDLER.PaymentCallbackSimulator)
		r.Post("/webhook", handler.HANDLER.PaymentWebhookSimulator)
	})

	handler.ROUTER.Route(helper.Version("payment"), func(r chi.Router) {
		r.Use(middleware.Auth(options.ENV.JWT.EXPIRED, options.REDIS))
		r.Post("/simulator", handler.HANDLER.SimulatorPayment)
		r.Post("/generate", handler.HANDLER.GeneratePayment)
		r.Get("/{id}/status", handler.HANDLER.CheckStatusPayment)
	})
}
