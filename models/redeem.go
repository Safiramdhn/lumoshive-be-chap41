package models

import "time"

type Redeem struct {
	Base
	UserID      int       `json:"user_id,omitempty" gorm:"type:integer" example:"1"`
	VoucherID   int       `gorm:"type:integer;index" json:"-" swaggerignore:"true"` // Foreign key to Voucher
	VoucherCode string    `gorm:"type:varchar(255)" json:"voucher_code,omitempty" example:"VCHR2024"`
	RedeemDate  time.Time `gorm:"type:timestamp" json:"redeem_date,omitempty" example:"2024-12-03T12:00:00Z"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID" json:"voucher,omitempty"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
