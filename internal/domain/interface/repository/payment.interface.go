package rinf

import (
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type IPaymentRepository interface {
	Find() *bun.SelectQuery
	FindOne() *bun.SelectQuery
	Create(model *model.PaymentModel) *bun.InsertQuery
	Update(model *model.PaymentModel) *bun.UpdateQuery
}
