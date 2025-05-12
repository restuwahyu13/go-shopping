package model

import (
	"encoding/json"
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type OrderItemModel struct {
	bun.BaseModel  `bun:"table:order_item"`
	OrderID        string          `bun:"order_id,notnull,unique" json:"order_id"`
	ProductItemID  string          `bun:"product_item_id,notnull" json:"product_item_id"`
	Qty            int64           `bun:"qty,notnull" json:"qty"`
	Amount         int64           `bun:"amount,notnull" json:"amount"`
	PromotionRules json.RawMessage `bun:"promotion_rules,nullzero,type:jsonb" json:"promotion_rules"`
	FreeProduct    []string        `bun:"free_product,nullzero" json:"free_product"`
	CreatedAt      time.Time       `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	CreatedBy      zero.String     `bun:"created_by,nullzero" json:"created_by"`
	UpdatedAt      zero.Time       `bun:"updated_at,nullzero" json:"updated_at"`
	UpdatedBy      zero.String     `bun:"updated_by,nullzero" json:"updated_by"`
	DeletedAt      zero.Time       `bun:"deleted_at,nullzero" json:"deleted_at"`
	DeletedBy      zero.String     `bun:"deleted_by,nullzero" json:"deleted_by"`

	/**
	 * ===========================================
	 *  DATABASE RELATION TABLE
	 * ===========================================
	 **/

	Order       *OrderModel       `bun:"rel:belongs-to,join:order_id=id" json:"order,omitempty"`
	ProductItem *ProductItemModel `bun:"rel:belongs-to,join:product_item_id=id" json:"product_item,omitempty"`
}
