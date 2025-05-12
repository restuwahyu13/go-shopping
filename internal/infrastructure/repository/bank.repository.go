package repo

import (
	"context"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type bankRepository struct {
	ctx   context.Context
	db    *bun.DB
	model *model.BankModel
}

func NewBankRepository(ctx context.Context, db *bun.DB) rinf.IBankRepository {
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
