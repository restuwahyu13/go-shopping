package repo

import (
	"context"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type (
	ICourierRepository interface {
		Find() *bun.SelectQuery
		FindOne() *bun.SelectQuery
		Create() *bun.InsertQuery
		Update() *bun.UpdateQuery
	}

	courierRepository struct {
		ctx   context.Context
		db    *bun.DB
		model *model.CourierModel
	}
)

func NewCourierRepository(ctx context.Context, db *bun.DB) ICourierRepository {
	return courierRepository{ctx: ctx, db: db, model: new(model.CourierModel)}
}

func (r courierRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r courierRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r courierRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r courierRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
