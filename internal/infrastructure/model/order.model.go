package model

import (
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type OrderModel struct {
	bun.BaseModel `bun:"table:order"`
	ID            string      `bun:"id,pk,default:uuid_generate_v4()" json:"id"`
	PaymentID     string      `bun:"payment_id,nullzero" json:"payment_id"`
	UserID        string      `bun:"user_id,notnull" json:"user_id"`
	CourierID     string      `bun:"courier_id,nullzero" json:"courier_id"`
	InvoiceNumber string      `bun:"invoice_number,notnull,unique" json:"invoice_number"`
	Status        string      `bun:"status,notnull" json:"status"`
	Paid          bool        `bun:"paid,nullzero,default:false" json:"paid"`
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

	Payment   *PaymentModel     `bun:"rel:belongs-to,join:payment_id=id" json:"payment,omitempty"`
	User      *UsersModel       `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	OrderItem *[]OrderItemModel `bun:"rel:has-many,join:id=order_id" json:"order_item,omitempty"`
}
