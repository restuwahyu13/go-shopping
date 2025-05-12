package model

import (
	"time"

	"github.com/guregu/null/v6/zero"
	"github.com/uptrace/bun"
)

type PaymentModel struct {
	bun.BaseModel `bun:"table:payment"`
	ID            string      `bun:"id,pk,default:uuid_generate_v4()" json:"id"`
	BankID        string      `bun:"bank_id,notnull" json:"bank_id"`
	UserID        string      `bun:"user_id,notnull" json:"user_id"`
	RequestID     string      `bun:"request_id,notnull,unique" json:"request_id"`
	Amount        int64       `bun:"amount,nullzero" json:"amount"`
	Status        string      `bun:"status,notnull" json:"status"`
	AccountNumber int64       `bun:"account_number,nullzero" json:"account_number"`
	Sender        string      `bun:"sender,nullzero" json:"sender"`
	VerifiedAt    zero.Time   `bun:"verified_at,nullzero" json:"verified_at"`
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

	Bank *BankModel  `bun:"rel:belongs-to,join:bank_id=id" json:"bank"`
	User *UsersModel `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}
