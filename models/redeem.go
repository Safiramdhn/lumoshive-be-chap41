package models

import "time"

type Redeem struct {
	Base
	UserID      int       `json:"user_id,omitempty" gorm:"type:integer"`
	VoucherID   int       `gorm:"type:integer;index" json:"-"` // Foreign key to Voucher
	VoucherCode string    `gorm:"type:varchar(255)" json:"voucher_code,,omitempty"`
	RedeemDate  time.Time `gorm:"type:timestamp" json:"redeem_date,omitempty"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID" json:"voucher"`
	User    User    `gorm:"foreignKey:UserID" json:"user"`
}
