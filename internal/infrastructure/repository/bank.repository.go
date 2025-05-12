package repo

import (
	"context"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

// func NewBankRepository(db *bun.DB)

type (
	IBankRepository interface {
		Find() *bun.SelectQuery
		FindOne() *bun.SelectQuery
		Create() *bun.InsertQuery
		Update() *bun.UpdateQuery
	}

	bankRepository struct {
		ctx   context.Context
		db    *bun.DB
		model *model.BankModel
	}
)

func NewBankRepository(ctx context.Context, db *bun.DB) IBankRepository {
	return bankRepository{ctx: ctx, db: db, model: new(model.BankModel)}
}

func (r bankRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r bankRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r bankRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r bankRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
