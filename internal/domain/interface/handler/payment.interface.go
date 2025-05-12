package hdinf

import "net/http"

type IPaymentHandler interface {
	PaymentCallbackSimulator(rw http.ResponseWriter, r *http.Request)
	PaymentWebhookSimulator(rw http.ResponseWriter, r *http.Request)
	PaymentSimulator(rw http.ResponseWriter, r *http.Request)
	PaymentStatus(rw http.ResponseWriter, r *http.Request)
}
