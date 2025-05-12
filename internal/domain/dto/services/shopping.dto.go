package sdto

type (
	// CheckoutOrderShoppingDTO struct {
	// 	ProductItemID string `json:"product_item_id" validate:"required,uuid4"`
	// 	Amount        int64  `json:"amount" validate:"required,number"`
	// 	Qty           int64  `json:"qty" validate:"required,number"`
	// 	Notes         string `json:"notes"`
	// 	Action        string `json:"action" validate:"required,oneof=checkout remove order"`
	// }

	CheckoutShoppingDTO struct {
		ProductItemID string `json:"product_item_id" validate:"required,uuid4"`
		CourierID     string `json:"courier_id" validate:"required,uuid4"`
		Amount        int64  `json:"amount" validate:"required,number"`
		Qty           int64  `json:"qty" validate:"required,number"`
		Notes         string `json:"notes"`
		Action        string `json:"action" validate:"required,oneof=checkout remove order"`
	}

	CheckoutShoppingCacheDTO struct {
		ProductItemID string `json:"product_item_id"`
		Qty           int64  `json:"qty"`
	}
)
