package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	sdto "restuwahyu13/shopping-cart/internal/domain/dto/services"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	sinf "restuwahyu13/shopping-cart/internal/domain/interface/service"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/model"
	repo "restuwahyu13/shopping-cart/internal/infrastructure/repository"

	"github.com/google/uuid"
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

func (s ShoppingService) Checkout(ctx context.Context, body sdto.CheckoutDTO) hopt.Response {
	var (
		res            hopt.Response = hopt.Response{}
		userId         string        = fmt.Sprintf("%v", ctx.Value("user_id"))
		message        string        = "Checkout product item success"
		sumTotalAmount int64         = 0
	)

	bankModel := model.BankModel{}
	courierModel := model.CourierModel{}
	paymentModel := new(model.PaymentModel)

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

	if helper.IsCheckoutOrder(body.Orders) {
		bankRepository := repo.NewBankRepository(ctx, s.DB)
		err := bankRepository.FindOne().Column("id").Where("deleted_at IS NULL and active = true AND id = ?", body.BankID).Scan(ctx, &bankModel.ID)
		if err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err == sql.ErrNoRows {
			res.StatCode = http.StatusNotFound
			res.ErrMsg = "Bank is not exist in our system"

			return res
		}

		courierRepository := repo.NewCourierRepository(ctx, s.DB)
		err = courierRepository.FindOne().Column("id").Where("deleted_at IS NULL and active = true AND id = ?", body.CourierID).Scan(ctx, &courierModel.ID)
		if err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err == sql.ErrNoRows {
			res.StatCode = http.StatusNotFound
			res.ErrMsg = "Courier is not exist in our system"

			return res
		}

		paymentModel.UserID = userId
		paymentModel.BankID = bankModel.ID
		paymentModel.RequestID = uuid.NewString()
		paymentModel.Status = cons.PENDING

		paymentResult, err := trx.NewInsert().Model(paymentModel).Returning("id").Exec(ctx, &paymentModel.ID)
		if err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		} else if rows, err := paymentResult.RowsAffected(); rows < 1 || err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			if err == nil {
				res.StatCode = http.StatusPreconditionFailed
				res.ErrMsg = "Failed to make a payment order"
			}

			return res
		}
	}

	for _, order := range body.Orders {
		productItemModel := model.ProductItemModel{}
		productItemRepository := repo.NewProductItemRepository(ctx, s.DB)

		err := productItemRepository.FindOne().Column("id", "sell_amount", "qty").Where("deleted_at IS NULL and ready = true AND id = ?", order.ProductItemID).Scan(ctx, &productItemModel.ID, &productItemModel.SellAmount, &productItemModel.Qty)
		if err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err == sql.ErrNoRows {
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
			}

			if err == sql.ErrNoRows {
				res.StatCode = http.StatusNotFound
				res.ErrMsg = "Product item is not exist in our system"

				return res
			}

			productItemRepository := repo.NewProductItemRepository(ctx, s.DB)
			productItemResult, err := productItemRepository.FindCheckoutProductItemByPromotionRules(trx, productConfigResult, body.Orders)

			if err != nil && err != sql.ErrNoRows {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res
			}

			if err == sql.ErrNoRows {
				res.StatCode = http.StatusNotFound
				res.ErrMsg = "Product item is not exist in our system"

				return res
			}

			orderModel := new(model.OrderModel)
			orderModel.UserID = userId
			orderModel.PaymentID = paymentModel.ID
			orderModel.CourierID = courierModel.ID
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
				orderItemModel.FreeProduct = []string{productItemResult.ProductItemID}
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

		isMember, err := s.RDS.SIsMember(keyCheckoutProduct, valueCheckoutProductItem)
		if err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if !isMember && order.Action != cons.CHECKOUT {
			res.StatCode = http.StatusUnprocessableEntity
			res.ErrMsg = "You must checkout product item before go to next step"

			return res
		}

		if isExist && isMember {
			resultPreviousCheckout, err := s.RDS.HGet(keyCheckoutProductCounter, order.ProductItemID)
			if err != nil {
				res.StatCode = http.StatusInternalServerError
				res.ErrMsg = err.Error()

				return res
			}
			previousCheckout := string(resultPreviousCheckout)

			if (order.Action == cons.CHECKOUT || order.Action == cons.ORDER) && order.Qty > 0 {
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

	if helper.IsCheckoutOrder(body.Orders) {
		resultUpdatePayment, err := trx.NewUpdate().Table("payment").Set("amount = ?", &sumTotalAmount).Where("id = ?", paymentModel.ID).Exec(ctx)
		if err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		} else if rows, err := resultUpdatePayment.RowsAffected(); rows < 1 || err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			if err == nil {
				res.StatCode = http.StatusPreconditionFailed
				res.ErrMsg = "Failed to make a payment order"
			}

			return res
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

func (s ShoppingService) CheckoutList(ctx context.Context) hopt.Response {
	var (
		res                hopt.Response         = hopt.Response{}
		checkout           sdto.CheckoutCacheDTO = sdto.CheckoutCacheDTO{}
		userId             any                   = ctx.Value("user_id")
		keyCheckoutProduct string                = fmt.Sprintf("CHECKOUT_PRODUCT:%s", userId)
	)

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
		if err := parser.Unmarshal([]byte(member), &checkout); err != nil {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err := productItemRepository.Find().Where("id = ?", checkout.ProductItemID).Scan(ctx, &productItemModel); err != nil && err != sql.ErrNoRows {
			res.StatCode = http.StatusInternalServerError
			res.ErrMsg = err.Error()

			return res
		}

		if err == sql.ErrNoRows {
			res.StatCode = http.StatusNotFound
			res.ErrMsg = "Checkout product item is not exist in our system"

			return res
		}

		productItemModel.Qty = checkout.Qty
		productItemModels = append(productItemModels, productItemModel)
	}

	res.StatCode = http.StatusOK
	res.Message = "Success"
	res.Data = productItemModels

	return res
}
