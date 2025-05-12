package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"

	"github.com/lib/pq"
	"github.com/uptrace/bun"
)

type ShoppingService struct {
	ENV *cdto.Environtment
	DB  *bun.DB
	RDS pinf.IRedis
}

func NewShoppingService(options ShoppingService) sinf.IShoppingService {
	return ShoppingService{
		ENV: options.ENV,
		DB:  options.DB,
		RDS: options.RDS,
	}
}

func (s ShoppingService) CreateCheckoutCartShopping(ctx context.Context, req hdto.Request[[]sdto.CheckoutShoppingDTO]) hopt.Response {
	res := hopt.Response{}

	trx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	defer func(err error) {
		if err != nil {
			if err := trx.Rollback(); err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()
			}
		} else {
			if err := trx.Commit(); err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

			}
		}
	}(err)

	userId := fmt.Sprintf("%v", ctx.Value("user_id"))
	message := "Checkout product item success"
	sumTotalAmount := int64(0)

	for _, order := range req.Body {
		productItemModel := model.ProductItemModel{}
		productItemRepository := repo.NewProductItemRepository(ctx, s.DB)

		err := productItemRepository.FindOne().Column("id", "sell_amount", "qty").Where("deleted_at IS NULL and ready = true AND id = ?", order.ProductItemID).Scan(ctx, &productItemModel.ID, &productItemModel.SellAmount, &productItemModel.Qty)
		if err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res

		} else if err == sql.ErrNoRows {
			res.StatCode = http.StatusNotFound
			res.ErrMsg = fmt.Sprintf("Product item ID %s not found", order.ProductItemID)

			return res
		}

		if order.Amount != productItemModel.SellAmount {
			res.StatCode = http.StatusUnprocessableEntity
			res.ErrMsg = "Invalid product item amount"

			return res
		}

		if order.Action == cons.ORDER && order.Qty > 0 {
			productConfigRepository := repo.NewProductConfigRepository(ctx, s.DB)
			productConfigResult, err := productConfigRepository.FindProductConfigByProductId(order.ProductItemID)

			if err != nil && err != sql.ErrNoRows {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res

			} else if err == sql.ErrNoRows {
				res.StatCode = http.StatusNotFound
				res.ErrMsg = "Product item is not exist in our system"

				return res
			}

			productItemRepository := repo.NewProductItemRepository(ctx, s.DB)
			productItemResult, err := productItemRepository.FindCheckoutProductItemByPromotionRules(trx, productConfigResult, req.Body)

			if err != nil && err != sql.ErrNoRows {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res

			} else if err == sql.ErrNoRows {
				res.StatCode = http.StatusNotFound
				res.ErrMsg = "Product item is not exist in our system"

				return res
			}

			orderModel := new(model.OrderModel)
			orderModel.UserID = userId
			orderModel.InvoiceNumber = fmt.Sprintf("INV-%s", helper.NewRandom().Numeric(8))
			orderModel.Status = cons.WAITING
			orderModel.Notes = order.Notes
			orderModel.OriginAmount = productItemResult.OriginAmount
			orderModel.DiscountAmount = productItemResult.DiscountAmount
			orderModel.TotalAmount = productItemResult.TotalAmount

			orderResult, err := trx.NewInsert().Model(orderModel).Returning("id").Exec(ctx, &orderModel.ID)
			if err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res

			} else if rows, err := orderResult.RowsAffected(); rows < 1 || err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				if err == nil {
					res.StatCode = http.StatusPreconditionFailed
					res.ErrMsg = "Failed to make a payment order"
				}

				return res
			}

			orderItemModel := new(model.OrderItemModel)
			orderItemModel.OrderID = orderModel.ID
			orderItemModel.ProductItemID = order.ProductItemID
			orderItemModel.Qty = order.Qty
			orderItemModel.Amount = order.Amount
			orderItemModel.PromotionRules = productConfigResult.PromotionRulesRaw

			if productItemResult.ProductItemID != "" {
				orderItemModel.FreeProduct = pq.Array([]string{productItemResult.ProductItemID})
			}

			orderItemResult, err := trx.NewInsert().Model(orderItemModel).Exec(ctx)
			if err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res

			} else if rows, err := orderItemResult.RowsAffected(); rows < 1 || err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				if err == nil {
					res.StatCode = http.StatusPreconditionFailed
					res.ErrMsg = "Failed to make a payment order"
				}

				return res
			}

			previousProductItemQty := productItemModel.Qty - order.Qty
			resultUpdateProductItem, err := trx.NewUpdate().Table("product_item").Set("qty = ?", &previousProductItemQty).Where("id = ?", productItemModel.ID).Exec(ctx)
			if err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res

			} else if rows, err := resultUpdateProductItem.RowsAffected(); rows < 1 || err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				if err == nil {
					res.StatCode = http.StatusPreconditionFailed
					res.ErrMsg = "Failed to make a payment order"
				}

				return res
			}

			sumTotalAmount += productItemResult.TotalAmount
		}

		keyCheckoutProduct := fmt.Sprintf("CHECKOUT_PRODUCT:%s", userId)
		keyCheckoutProductCounter := fmt.Sprintf("CHECKOUT_PRODUCT_COUNTER:%s", userId)
		valueCheckoutProductItem := fmt.Sprintf(`{"product_item_id":"%s","qty":%d}`, order.ProductItemID, order.Qty)

		isExist, err := s.RDS.HExists(keyCheckoutProductCounter, order.ProductItemID)
		if err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if isExist {
			resultPreviousCheckout, err := s.RDS.HGet(keyCheckoutProductCounter, order.ProductItemID)
			if err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res
			}
			previousCheckout := string(resultPreviousCheckout)

			if (order.Action == cons.CHECKOUT || order.Action == cons.ORDER) && order.Qty > 0 {
				fmt.Println("keyCheckoutProduct - dalam", previousCheckout)

				if _, err := s.RDS.SRem(keyCheckoutProduct, previousCheckout); err != nil {
					res.StatCode = http.StatusInternalServerError
					res.ErrMsg = err.Error()

					return res
				}

				if _, err := s.RDS.HDel(keyCheckoutProductCounter, order.ProductItemID); err != nil {
					res.StatCode = http.StatusInternalServerError
					res.ErrMsg = err.Error()

					return res
				}

			} else if order.Action == cons.REMOVE && order.Qty < 2 {
				if _, err := s.RDS.SRem(keyCheckoutProduct, previousCheckout); err != nil {
					res.StatCode = http.StatusInternalServerError
					res.ErrMsg = err.Error()

					return res
				}

				if _, err := s.RDS.HDel(keyCheckoutProductCounter, order.ProductItemID); err != nil {
					res.StatCode = http.StatusInternalServerError
					res.ErrMsg = err.Error()

					return res
				}

			} else {
				res.StatCode = http.StatusPreconditionFailed
				res.ErrMsg = fmt.Sprintf("%s product item failed", order.Action)

				return res
			}
		}

		if order.Action != cons.REMOVE && order.Action != cons.ORDER {
			if err := s.RDS.HSet(keyCheckoutProductCounter, order.ProductItemID, valueCheckoutProductItem); err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res
			}

			if _, err := s.RDS.SAdd(keyCheckoutProduct, valueCheckoutProductItem); err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res
			}
		}

		if order.Action == cons.ORDER {
			message = "Order product item success, please make a payment"
		} else if order.Action == cons.REMOVE {
			message = "Remove product item success"
		}
	}

	keyCheckoutProduct := fmt.Sprintf("CHECKOUT_PRODUCT:%s", userId)
	members, err := s.RDS.SMembers(keyCheckoutProduct)
	if err != nil {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res
	}

	res.StatCode = http.StatusOK
	res.Message = message
	res.Data = members

	return res
}

func (s ShoppingService) ListCheckoutCartShopping(ctx context.Context) hopt.Response {
	userId := ctx.Value("user_id")
	keyCheckoutProduct := fmt.Sprintf("CHECKOUT_PRODUCT:%s", userId)

	res := hopt.Response{}
	req := hdto.Request[sdto.CheckoutShoppingCacheDTO]{}

	members, err := s.RDS.SMembers(keyCheckoutProduct)
	if err != nil {
		res.StatCode = http.StatusInternalServerError
		res.ErrMsg = err.Error()

		return res
	}

	parser := helper.NewParser()
	productItemModel := model.ProductItemModel{}
	productItemModels := []model.ProductItemModel{}
	productItemRepository := repo.NewProductItemRepository(ctx, s.DB)

	for _, member := range members {
		if err := parser.Unmarshal([]byte(member), &req.Body); err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err := productItemRepository.Find().Where("id = ?", req.Body.ProductItemID).Scan(ctx, &productItemModel); err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err == sql.ErrNoRows {
			res.StatCode = http.StatusNotFound
			res.ErrMsg = "Checkout product item is not exist in our system"

			return res
		}

		productItemModel.Qty = req.Body.Qty
		productItemModels = append(productItemModels, productItemModel)
	}

	res.StatCode = http.StatusOK
	res.Message = "Success"
	res.Data = productItemModels

	return res
}
