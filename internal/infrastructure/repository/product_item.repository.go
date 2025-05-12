package repo

import (
	"context"
	"errors"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	rinf "restuwahyu13/shopping-cart/internal/domain/interface/repository"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	ropt "restuwahyu13/shopping-cart/internal/domain/output/repository"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"

	"github.com/uptrace/bun"
)

type productItemRepository struct {
	ctx   context.Context
	db    *bun.DB
	rds   pinf.IRedis
	model *model.ProductItemModel
}

func NewProductItemRepository(ctx context.Context, db *bun.DB) rinf.IProductItemRepository {
	return productItemRepository{ctx: ctx, db: db, model: new(model.ProductItemModel)}
}

func (r productItemRepository) Find() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productItemRepository) FindOne() *bun.SelectQuery {
	return r.db.NewSelect().Model(r.model)
}

func (r productItemRepository) Create() *bun.InsertQuery {
	return r.db.NewInsert().Model(r.model)
}

func (r productItemRepository) Update() *bun.UpdateQuery {
	return r.db.NewUpdate().Model(r.model)
}

func (r productItemRepository) FindCheckoutProductItemByPromotionRules(trx bun.Tx, promotionConfig *ropt.FindProductConfigByProductId, checkoutShoppings []sdto.CheckoutShoppingDTO) (*ropt.FindCheckoutProductItemPromotionRules, error) {
	var (
		productItemModel *model.ProductItemModel                     = new(model.ProductItemModel)
		res              *ropt.FindCheckoutProductItemPromotionRules = new(ropt.FindCheckoutProductItemPromotionRules)
		parser           hinf.IParser                                = helper.NewParser()
		rules            *hopt.PromotionRules                        = helper.PromotionRules(promotionConfig.PromotionRules)
	)

	if rules != nil && promotionConfig.ProductItemID != "" {
		for _, product := range checkoutShoppings {
			err := trx.NewSelect().Model(productItemModel).Model(productItemModel).
				Column("id", "sub_brand", "category", "sell_amount").
				Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND  id = ?)", promotionConfig.ProductItemID).
				Scan(r.ctx, &productItemModel.ID, &productItemModel.SubBrand, &productItemModel.Category, &productItemModel.SellAmount)

			if err != nil {
				return nil, err
			}

			if rules.Type == cons.BUY_ONE_GET_ONE {
				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.TotalAmount = int64(calculateAmount)

				err := trx.NewSelect().Model(productItemModel).Column("id").
					Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id = ? AND id != ?)", rules.OptionOne, promotionConfig.ProductItemID).
					Scan(r.ctx, &productItemModel.ID)

				if err != nil {
					return nil, err
				}

				res.PreviousProductID = promotionConfig.ProductItemID
				res.ProductItemID = productItemModel.ID
				return res, nil
			} else if rules.Type == cons.BUY_GET_OTHER_PRODUCT {
				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.TotalAmount = int64(calculateAmount)

				err := trx.NewSelect().Model(productItemModel).Column("id").
					Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id = ? AND id != ?)", rules.OptionOne, promotionConfig.ProductItemID).
					Scan(r.ctx, &productItemModel.ID)
				if err != nil {
					return nil, err
				}

				res.PreviousProductID = promotionConfig.ProductItemID
				res.ProductItemID = productItemModel.ID
				return res, nil
			} else if rules.Type == cons.BUY_GET_DISCOUNT_PERCENTAGE {
				parseInt, err := parser.ToInt(rules.OptionOne)
				if err != nil {
					return nil, err
				}

				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
				totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.PERCENTAGE)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.DiscountAmount = int64(discountAmount)
				res.TotalAmount = int64(totalAmount)

				return res, nil
			} else if rules.Type == cons.BUY_GET_DISCOUNT_PERCENTAGE_WITH_PRODUCT {
				if rules.OptionThree == cons.SAME_BRAND_SIMILAR_CATEGORY {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.PERCENTAGE)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0  AND id != ? AND sub_brand = ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionThree == cons.SAME_BRAND_ANY_CATEGORY {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.PERCENTAGE)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0  AND id != ? AND sub_brand = ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionThree == cons.ANY_BRAND_ANY_CATEGORY {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.PERCENTAGE)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id != ? AND sub_brand != ? AND category != ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionThree == cons.ANY_BRAND_SAME_PRODUCT {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.PERCENTAGE)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id != ? AND sub_brand != ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil
				}

			} else if rules.Type == cons.BUY_GET_DISCOUNT_FIXED {
				parseInt, err := parser.ToInt(rules.OptionOne)
				if err != nil {
					return nil, err
				}

				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
				totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.DiscountAmount = int64(discountAmount)
				res.TotalAmount = int64(totalAmount)

				return res, nil
			} else if rules.Type == cons.BUY_GET_DISCOUNT_FIXED_WITH_PRODUCT {
				parseInt, err := parser.ToInt(rules.OptionOne)
				if err != nil {
					return nil, err
				}

				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
				totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.DiscountAmount = int64(discountAmount)
				res.TotalAmount = int64(totalAmount)

				if rules.OptionFour == cons.SAME_BRAND_SIMILAR_CATEGORY {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0  AND id != ? AND sub_brand = ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionFour == cons.SAME_BRAND_ANY_CATEGORY {

					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0  AND id != ? AND sub_brand = ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionFour == cons.ANY_BRAND_ANY_CATEGORY {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id != ? AND sub_brand != ? AND category != ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil

				} else if rules.OptionFour == cons.ANY_BRAND_SAME_PRODUCT {
					parseInt, err := parser.ToInt(rules.OptionOne)
					if err != nil {
						return nil, err
					}

					calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
					totalAmount, discountAmount := helper.PromotionDiscount(calculateAmount, int64(parseInt), cons.FIXED)

					res.Key = rules.Key
					res.Type = rules.Type
					res.OriginAmount = int64(calculateAmount)
					res.DiscountAmount = int64(discountAmount)
					res.TotalAmount = int64(totalAmount)

					err = trx.NewSelect().Model(productItemModel).Column("id").
						Where("(deleted_at IS NULL AND ready = true AND qty > 0 AND id != ? AND sub_brand != ? AND category = ?)", promotionConfig.ProductItemID, productItemModel.SubBrand, productItemModel.Category).
						Order("created_at DESC").Scan(r.ctx, &productItemModel.ID)

					if err != nil {
						return nil, err
					}

					res.PreviousProductID = promotionConfig.ProductItemID
					res.ProductItemID = productItemModel.ID
					return res, nil
				}
			} else if rules.Type == cons.BUY_GET_HALF_PRICE {
				calculateAmount := float64(productItemModel.SellAmount) * float64(product.Qty)
				totalAmount := calculateAmount - float64(productItemModel.SellAmount)

				res.Key = rules.Key
				res.Type = rules.Type
				res.OriginAmount = int64(calculateAmount)
				res.TotalAmount = int64(totalAmount)

				return res, nil
			}

			return nil, errors.New("Failed to use promotion rules, for checkout product item")
		}
	}

	return nil, errors.New("Failed to use promotion rules, for checkout product item")
}
