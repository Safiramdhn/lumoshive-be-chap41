package models

import "time"

type Usage struct {
	Base
	UserID            int       `json:"user_id,omitempty" gorm:"type:integer" example:"1"`
	VoucherID         int       `gorm:"type:integer;index" json:"-" swaggerignore:"true"` // Foreign key to Voucher
	VoucherCode       string    `gorm:"type:varchar(255)" json:"voucher_code,omitempty" example:"VCHR2024"`
	UsageDate         time.Time `gorm:"type:timestamp" json:"usage_date,omitempty" example:"2024-12-03T12:00:00Z"`
	TransactionAmount float64   `gorm:"type:decimal(10,2)" json:"transaction_amount,omitempty" example:"100.50"`
	BenefitAmount     float64   `gorm:"type:decimal(10,2)" json:"benefit_amount,omitempty" example:"10.00"`

	// Relationships
	Voucher Voucher `gorm:"foreignKey:VoucherID" json:"voucher,omitempty"`
	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type UsageDTO struct {
	UserID       int        `json:"user_id" example:"1"`
	VoucherInput VoucherDTO `json:"voucher_input"`
}
