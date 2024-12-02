package models

import (
	"time"
)

type Voucher struct {
	Base
	VoucherName     string    `gorm:"type:varchar(255)" json:"voucher_name,omitempty"`
	VoucherCode     string    `gorm:"type:varchar(255);uniqueIndex" json:"voucher_code,omitempty"`
	VoucherType     string    `gorm:"type:voucher_type" json:"voucher_type,omitempty"`
	Description     string    `gorm:"type:text" json:"description,omitempty"`
	VoucherCategory string    `gorm:"type:voucher_category" json:"voucher_category,omitempty"`
	DiscountAmount  float64   `gorm:"type:decimal(10,2)" json:"discount_amount,omitempty"`
	MinPurchase     float64   `gorm:"type:decimal(10,2)" json:"min_purchase,omitempty"`
	PaymentMethod   string    `gorm:"type:varchar(255)" json:"payment_method,omitempty"`
	StartDate       time.Time `gorm:"type:date" json:"start_date,omitempty"`
	EndDate         time.Time `gorm:"type:date" json:"end_date,omitempty"`
	ApplicableAreas string    `gorm:"type:jsonb" json:"applicable_areas,omitempty"`
	Quantity        int       `gorm:"type:int" json:"quantity,omitempty"`

	// Relationships
	Redeems []Redeem `gorm:"foreignKey:VoucherID" json:"redeems,omitempty"`
	Usages  []Usage  `gorm:"foreignKey:VoucherID" json:"usages,omitempty"`

	MinRatePoint int `gorm:"type:integer" json:"min_rate_point,omitempty"`
}

type VoucherWithStatus struct {
	Voucher
	IsActive bool `json:"is_active,omitempty"`
}

func (v *Voucher) IsActive() bool {
	now := time.Now()

	if now.After(v.StartDate) && now.Before(v.EndDate.Add(24*time.Hour)) {
		return true
	}
	return false
}

type VoucherDTO struct {
	VoucherCode             string    `json:"voucher_code"`
	TotalTransaction        float64   `json:"total_transactions"`
	TotalShippingCost       float64   `json:"total_shipping_cost"`
	TransactionDate         string    `json:"transaction_date"`
	FormatedTransactionDate time.Time `json:"-"`
	PaymentMethod           string    `json:"payment_method"`
	Area                    string    `json:"area"`
}

type ValidateVoucherResponse struct {
	TotalTransaction  float64 `json:"total_transaction"`
	TotalShippingCost float64 `json:"total_shipping_cost"`
	VoucherStatus     string  `json:"voucher_status"`
	BenefitAmount     float64 `json:"benefit_amount"`
	VoucherCode       string  `json:"-"`
	VoucherID         int     `json:"-"`
}
