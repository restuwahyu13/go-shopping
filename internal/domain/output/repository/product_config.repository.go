package ropt

import (
	"encoding/json"
	hdto "restuwahyu13/shopping-cart/internal/domain/dto/helper"
)

type FindProductConfigByProductId struct {
	ID                string                   `json:"id"`
	Name              string                   `json:"name"`
	Active            bool                     `json:"active"`
	PromotionRulesRaw json.RawMessage          `json:"promotion_rules_raw"`
	PromotionRules    []hdto.PromotionRulesDTO `json:"promotion_rules"`
	ProductID         string                   `json:"product_id"`
	ProductItemID     string                   `json:"product_item_id"`
	MinAmount         int64                    `json:"min_amount"`
	MaxAmount         int64                    `json:"max_amount"`
	ExpiredAt         string                   `json:"expired_at"`
	DeletedAt         any                      `json:"deleted_at"`
}
