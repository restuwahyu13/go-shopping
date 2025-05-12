package repo

import (
	"context"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type (
	IPaymentRepository interface {
		Find() *bun.SelectQuery
		FindOne() *bun.SelectQuery
		Create(model *model.PaymentModel) *bun.InsertQuery
		Update(model *model.PaymentModel) *bun.UpdateQuery
	}

	paymentRepository struct {
		ctx   context.Context
		db    *bun.DB
		model *model.PaymentModel
	}
)

func NewPaymentRepository(ctx context.Context, db *bun.DB) IPaymentRepository {
	return paymentRepository{ctx: ctx, db: db, model: new(model.PaymentModel)}
}

func (r paymentRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r paymentRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r paymentRepository) Create(model *model.PaymentModel) *bun.InsertQuery {
	return r.db.NewInsert().Model(model)
}

func (r paymentRepository) Update(model *model.PaymentModel) *bun.UpdateQuery {
	return r.db.NewUpdate().Model(model)
}
