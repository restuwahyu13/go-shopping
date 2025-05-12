package hdinf

import "net/http"

type IPaymentHandler interface {
	CallbackSimulatorPayment(rw http.ResponseWriter, r *http.Request)
	SimulatorPayment(rw http.ResponseWriter, r *http.Request)
	GeneratePayment(rw http.ResponseWriter, r *http.Request)
	CheckStatusPayment(rw http.ResponseWriter, r *http.Request)
}
