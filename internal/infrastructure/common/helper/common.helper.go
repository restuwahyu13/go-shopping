package helper

import (
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
)

func IsCheckoutOrder(items []sdto.CheckoutShoppingDTO) bool {
	for _, item := range items {
		if item.Action == cons.ORDER && item.Qty > 0 {
			return true
		}
	}
	return false
}
