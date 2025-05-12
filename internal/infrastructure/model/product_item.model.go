package model

import (
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type ProductItemModel struct {
	bun.BaseModel `bun:"product_item"`
	ID            string      `bun:"id,pk,default:uuid_generate_v4()" json:"id"`
	ProductID     string      `bun:"product_id,notnull" json:"product_id"`
	Name          string      `bun:"name,notnull" json:"name"`
	SKU           string      `bun:"sku,notnull,unique" json:"sku"`
	Qty           int64       `bun:"qty,notnull" json:"qty"`
	SubBrand      string      `bun:"sub_brand,notnull" json:"sub_brand"`
	Variant       string      `bun:"variant,notnull" json:"variant"`
	Category      string      `bun:"category,notnull" json:"category"`
	SerialNumber  string      `bun:"serial_number,notnull,unique" json:"serial_number"`
	Unit          string      `bun:"unit,notnull" json:"unit"`
	BuyAmount     int64       `bun:"buy_amount,notnull" json:"buy_amount"`
	SellAmount    int64       `bun:"sell_amount,notnull" json:"sell_amount"`
	Ready         bool        `bun:"ready,notnull,default:false" json:"ready"`
	Description   string      `bun:"description,nullzero" json:"description"`
	CreatedAt     time.Time   `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	CreatedBy     zero.String `bun:"created_by,nullzero" json:"created_by"`
	UpdatedAt     zero.Time   `bun:"updated_at,nullzero" json:"updated_at"`
	UpdatedBy     zero.String `bun:"updated_by,nullzero" json:"updated_by"`
	DeletedAt     zero.Time   `bun:"deleted_at,nullzero" json:"deleted_at"`
	DeletedBy     zero.String `bun:"deleted_by,nullzero" json:"deleted_by"`

	/**
	 * ===========================================
	 *  DATABASE RELATION TABLE
	 * ===========================================
	 **/

	Product *ProductModel `bun:"rel:belongs-to,join:product_id=id" json:"product,omitempty"`
}
