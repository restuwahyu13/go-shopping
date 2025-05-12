package model

import (
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type ProductModel struct {
	bun.BaseModel `bun:"table:product"`
	ID            string      `bun:"id,pk,default:uuid_generate_v4()" json:"id"`
	Brand         string      `bun:"brand,notnull,unique" json:"brand"`
	Code          string      `bun:"code,notnull" json:"code"`
	Active        bool        `bun:"active,notnull" json:"active"`
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

	ProductItem *[]ProductModel `bun:"rel:hash-many,join:id=product_id" json:"product_items,omitempty"`
}
