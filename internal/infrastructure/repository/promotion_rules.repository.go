package repo

import (
	"context"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type promotionRulesRepository struct {
	ctx   context.Context
	db    *bun.DB
	model *model.PromotionRulesModel
}

func NewPromotionRulesRepository(ctx context.Context, db *bun.DB) rinf.IPromotionRulesRepository {
	return promotionRulesRepository{ctx: ctx, db: db, model: new(model.PromotionRulesModel)}
}

func (r promotionRulesRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r promotionRulesRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r promotionRulesRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r promotionRulesRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}
