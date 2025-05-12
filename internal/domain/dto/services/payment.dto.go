package sdto

type (
	CallbackSimulatorPaymentDTO struct {
		TransactionID  string `json:"transaction_id" validate:"required,uuid4"`
		IdempotencyKey string `json:"idempotency_key" validate:"required,uuid4"`
		Bank           string `json:"bank" validate:"required,oneof=bca mandiri bni bri jago qris gopay ovo dana shopee"`
		Method         string `json:"method" validate:"required,oneof=va ewallet transfer qris"`
		Amount         int64  `json:"amount" validate:"required,number"`
		Status         string `json:"status" validate:"required,oneof=succeed failed"`
	}

	WebhookSimulatorPaymentDTO struct {
		TransactionID  string `json:"transaction_id" validate:"required,uuid4"`
		IdempotencyKey string `json:"idempotency_key" validate:"required,uuid4"`
		Bank           string `json:"bank" validate:"required,oneof=bca mandiri bni bri jago qris gopay ovo dana shopee"`
		Method         string `json:"method" validate:"required,oneof=va ewallet transfer qris"`
		Amount         int64  `json:"amount" validate:"required,number"`
		Status         string `json:"status" validate:"required,oneof=success failed, refund"`
	}

	SimulatorPaymentDTO struct {
		PaymentCode string `json:"payment_code" validate:"required,number"`
		Bank        string `json:"bank" validate:"required,oneof=bca mandiri bni bri jago qris gopay ovo dana shopee"`
		Method      string `json:"method" validate:"required,oneof=va ewallet transfer qris"`
		Amount      int64  `json:"amount" validate:"required,number"`
	}

	GeneratePaymentDTO struct {
		Bank   string `json:"bank" validate:"required,oneof=bca mandiri bni bri jago qris gopay ovo dana shopee"`
		Method string `json:"method" validate:"required,oneof=va ewallet transfer qris"`
		Amount int64  `json:"amount" validate:"required,number"`
	}

	CheckStatusPaymentDTO struct {
		ID string `json:"id" validate:"required,uuid4"`
	}
)
