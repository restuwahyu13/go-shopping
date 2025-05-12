package repo

import (
	"context"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type productRepository struct {
	ctx   context.Context
	db    *bun.DB
	model *model.ProductModel
}

func NewProductRepository(ctx context.Context, db *bun.DB) rinf.IProductRepository {
	return productRepository{ctx: ctx, db: db, model: new(model.ProductModel)}
}

func (r productRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r productRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
