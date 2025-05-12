package repo

import (
	"context"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	ropt "restuwahyu13/shopping-cart/internal/domain/output/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type (
	IProductConfigRepository interface {
		Find() *bun.SelectQuery
		FindOne() *bun.SelectQuery
		Create() *bun.InsertQuery
		Update() *bun.UpdateQuery
		FindProductConfigByProductId(productId string) (*ropt.FindProductConfigByProductId, error)
	}

	productConfigRepository struct {
		ctx   context.Context
		db    *bun.DB
		model *model.ProductConfig
	}
)

func NewProductConfigRepository(ctx context.Context, db *bun.DB) IProductConfigRepository {
	return productConfigRepository{ctx: ctx, db: db, model: new(model.ProductConfig)}
}

func (r productConfigRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productConfigRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productConfigRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r productConfigRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}

func (r productConfigRepository) FindProductConfigByProductId(productId string) (*ropt.FindProductConfigByProductId, error) {
	productConfigModel := new(model.ProductConfig)

	err := r.Find().Where("deleted_at IS NULL AND active = true AND (product_id = ? OR product_item_id = ?)", productId, productId).Scan(r.ctx, productConfigModel)
	if err != nil {
		return nil, err
	}

	promotionRulesByte, err := productConfigModel.PromotionRules.MarshalJSON()
	if err != nil {
		return nil, err
	}

	parser := helper.NewParser()
	promotionRules := []hdto.PromotionRulesDTO{}

	if err := parser.Unmarshal(promotionRulesByte, &promotionRules); err != nil {
		return nil, err
	}

	res := new(ropt.FindProductConfigByProductId)
	res.ID = productConfigModel.ID
	res.Name = productConfigModel.Name
	res.PromotionRulesRaw = productConfigModel.PromotionRules
	res.PromotionRules = promotionRules
	res.Active = productConfigModel.Active
	res.ProductID = productConfigModel.ProductID
	res.ProductItemID = productConfigModel.ProductItemID
	res.MinAmount = productConfigModel.MinAmount
	res.MaxAmount = productConfigModel.MaxAmount
	res.ExpiredAt = productConfigModel.ExpiredAt
	res.DeletedAt = productConfigModel.DeletedAt

	return res, nil
}
