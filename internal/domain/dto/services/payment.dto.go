package sdto

type (
	PaymentDTO struct {
		AccountNumber string `json:"account_number" validate:"required,number"`
		Sender        string `json:"sender" validate:"required"`
		Amount        int64  `json:"amount" validate:"required,number"`
	}

	PaymentStatusDTO struct {
		ID string `json:"id" validate:"required,uuid4"`
	}
)
