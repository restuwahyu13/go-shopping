package popt

import (
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"

	"github.com/go-chi/chi/v5"
)

type GracefulConfig struct {
	HANDLER *chi.Mux
	ENV     *cdto.Environtment
}
