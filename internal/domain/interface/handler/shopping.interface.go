package hdinf

import (
	"net/http"
)

type IShoppingHandler interface {
	Checkout(rw http.ResponseWriter, r *http.Request)
	CheckoutList(rw http.ResponseWriter, r *http.Request)
}
