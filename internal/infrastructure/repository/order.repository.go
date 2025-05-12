package repo

import (
	"context"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type orderRepository struct {
	ctx   context.Context
	db    *bun.DB
	model *model.OrderModel
}

func NewOrderRepository(ctx context.Context, db *bun.DB) rinf.IOrderRepository {
	return orderRepository{ctx: ctx, db: db, model: new(model.OrderModel)}
}

func (r orderRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r orderRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r orderRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r orderRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
