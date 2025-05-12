package hdinf

import (
	"net/http"
)

type IShoppingHandler interface {
	CreateCheckoutCartShopping(rw http.ResponseWriter, r *http.Request)
	ListCheckoutCartShopping(rw http.ResponseWriter, r *http.Request)
}
