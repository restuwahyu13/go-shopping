package repo

import (
	"context"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type orderItemRepository struct {
	ctx   context.Context
	db    *bun.DB
	model *model.OrderItemModel
}

func NewOrderItemRepository(ctx context.Context, db *bun.DB) rinf.IOrderItemRepository {
	return orderItemRepository{ctx: ctx, db: db, model: new(model.OrderItemModel)}
}

func (r orderItemRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r orderItemRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r orderItemRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r orderItemRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
