package models

import (
	"time"
)

type Voucher struct {
	Base
	VoucherName     string    `gorm:"type:varchar(255)" json:"voucher_name,omitempty" example:"Holiday Discount"` // Example included
	VoucherCode     string    `gorm:"type:varchar(255);uniqueIndex" json:"voucher_code,omitempty" example:"HOLIDAY2024"`
	VoucherType     string    `gorm:"type:voucher_type" json:"voucher_type,omitempty" example:"Percentage"`
	Description     string    `gorm:"type:text" json:"description,omitempty" example:"10% off on all purchases"`
	VoucherCategory string    `gorm:"type:voucher_category" json:"voucher_category,omitempty" example:"Seasonal"`
	DiscountAmount  float64   `gorm:"type:decimal(10,2)" json:"discount_amount,omitempty" example:"10.50"`
	MinPurchase     float64   `gorm:"type:decimal(10,2)" json:"min_purchase,omitempty" example:"100.00"`
	PaymentMethod   string    `gorm:"type:varchar(255)" json:"payment_method,omitempty" example:"Credit Card"`
	StartDate       time.Time `gorm:"type:date" json:"start_date,omitempty" example:"2024-12-01"`
	EndDate         time.Time `gorm:"type:date" json:"end_date,omitempty" example:"2024-12-31"`
	ApplicableAreas string    `gorm:"type:jsonb" json:"applicable_areas,omitempty" example:"[\"Jakarta\", \"Bandung\"]"`
	Quantity        int       `gorm:"type:int" json:"quantity,omitempty" example:"500"`

	// Relationships
	Redeems []Redeem `gorm:"foreignKey:VoucherID" json:"redeems,omitempty"`
	Usages  []Usage  `gorm:"foreignKey:VoucherID" json:"usages,omitempty"`

	MinRatePoint int `gorm:"type:integer" json:"min_rate_point,omitempty" example:"50"`
}

type VoucherWithStatus struct {
	Voucher
	IsActive bool `json:"is_active,omitempty" example:"true"`
}

func (v *Voucher) IsActive() bool {
	now := time.Now()

	if now.After(v.StartDate) && now.Before(v.EndDate.Add(24*time.Hour)) {
		return true
	}
	return false
}

type VoucherDTO struct {
	VoucherCode             string    `json:"voucher_code" example:"HOLIDAY2024"`
	TotalTransaction        float64   `json:"total_transactions" example:"200.00"`
	TotalShippingCost       float64   `json:"total_shipping_cost" example:"10.00"`
	TransactionDate         string    `json:"transaction_date" example:"2024-12-03"`
	FormatedTransactionDate time.Time `json:"-"` // Internal use only
	PaymentMethod           string    `json:"payment_method" example:"Credit Card"`
	Area                    string    `json:"area" example:"Jakarta"`
}

type ValidateVoucherResponse struct {
	TotalTransaction  float64 `json:"total_transaction" example:"200.00"`
	TotalShippingCost float64 `json:"total_shipping_cost" example:"10.00"`
	VoucherStatus     string  `json:"voucher_status" example:"Valid"`
	BenefitAmount     float64 `json:"benefit_amount" example:"20.00"`
	VoucherCode       string  `json:"-"` // Internal use only
	VoucherID         int     `json:"-"` // Internal use only
}
