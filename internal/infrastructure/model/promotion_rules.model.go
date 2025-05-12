package model

import (
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type PromotionRulesModel struct {
	bun.BaseModel `bun:"table:promotion_rules"`
	ID            string      `bun:"bun:id,pk,default:uuid_generate_v4()" json:"id"`
	Name          string      `bun:"name,notnull,unique" json:"name"`
	Code          string      `bun:"code,notnull,unique" json:"code"`
	Aritmetic     string      `bun:"aritmetic,nullzero" json:"aritmetic"`
	CreatedAt     time.Time   `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	CreatedBy     zero.String `bun:"created_by,nullzero" json:"created_by"`
	UpdatedAt     zero.Time   `bun:"updated_at,nullzero" json:"updated_at"`
	UpdatedBy     zero.String `bun:"updated_by,nullzero" json:"updated_by"`
	DeletedAt     zero.Time   `bun:"deleted_at,nullzero" json:"deleted_at"`
	DeletedBy     zero.String `bun:"deleted_by,nullzero" json:"deleted_by"`
}
