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
)
