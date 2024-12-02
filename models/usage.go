package models

import "time"

type Usage struct {
	Base
	UserID            int       `json:"user_id,omitempty" gorm:"type:integer"`
	VoucherID         int       `gorm:"type:integer;index" json:"-"` // Foreign key to Voucher
	VoucherCode       string    `gorm:"type:varchar(255)" json:"voucher_code,omitempty"`
	UsageDate         time.Time `gorm:"type:timestamp" json:"usage_date,omitempty"`
	TransactionAmount float64   `gorm:"type:decimal(10,2)" json:"transaction_amount,omitempty"`
	BenefitAmount     float64   `gorm:"type:decimal(10,2)" json:"benefit_amount,omitempty"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID" json:"voucher,omitempty"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type UsageDTO struct {
	UserID       int        `json:"user_id"`
	VoucherInput VoucherDTO `json:"voucher_input"`
}
