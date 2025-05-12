package model

import (
	"encoding/json"
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type ProductConfig struct {
	bun.BaseModel  `bun:"table:product_config"`
	ID             string          `bun:"id,pk,default:uuid_generate_v4()" json:"id"`
	Name           string          `bun:"name,notnull" json:"name"`
	Active         bool            `bun:"active,notnull" json:"active"`
	PromotionRules json.RawMessage `bun:"promotion_rules,type:jsonb" json:"promotion_rules"`
	ProductID      string          `bun:"product_id,nullzero" json:"product_id"`
	ProductItemID  string          `bun:"product_item_id,nullzero" json:"product_item_id"`
	MinAmount      int64           `bun:"min_amount,nullzero,default:0" json:"min_amount"`
	MaxAmount      int64           `bun:"max_amount,nullzero,default:0" json:"max_amount"`
	ExpiredAt      string          `bun:"expired_at,nullzero" json:"expired_at"`
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

	Product     *ProductModel     `bun:"rel:belongs-to,join:product_id=id" json:"product,omitempty"`
	ProductItem *ProductItemModel `bun:"rel:belongs-to,join:product_item_id=id" json:"product_item,omitempty"`
}
