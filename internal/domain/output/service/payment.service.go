package sopt

type (
	GeneratePayment struct {
		RequestID   string `json:"request_id"`
		PaymentCode string `json:"payment_code"`
		Bank        string `json:"bank"`
		Method      string `json:"method"`
		Amount      int64  `json:"amount"`
		ExpiredAt   string `json:"expired_at"`
		CreatedAt   string `json:"created_at"`
	}

	CallbackSimulatorPayment struct {
		TransactionID  string `json:"transaction_id" validate:"required,uuid4"`
		IdempotencyKey string `json:"idempotency_key" validate:"required,uuid4"`
		Bank           string `json:"bank" validate:"required,oneof=bca mandiri bni bri jago qris gopay ovo dana shopee"`
		Method         string `json:"method" validate:"required,oneof=va ewallet transfer qris"`
		Amount         int64  `json:"amount" validate:"required,number"`
		Status         string `json:"status" validate:"required,oneof=success failed, refund"`
	}
)
