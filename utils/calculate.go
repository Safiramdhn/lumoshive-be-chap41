package utils

import (
	"math"
)

func CalculateBenefit(totalTransaction, TotalShippingCost, DiscountAmount float64, category string) float64 {
	switch category {
	case "discount":
		return math.Round(totalTransaction*((100.00-DiscountAmount)/100.00)*100) / 100
	case "free_shipping":
		return math.Round(TotalShippingCost*((100.00-DiscountAmount)/100.00)*100) / 100
	default:
		return 0
	}
}
