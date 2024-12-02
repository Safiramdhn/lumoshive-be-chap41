package utils

import (
	"encoding/json"
	"errors"
	"log"
	"lumoshive-be-chap41/models"

	"time"
)

func ValidateVoucher(voucherValidateInput models.VoucherDTO, voucher models.Voucher) (models.ValidateVoucherResponse, error) {
	var result models.ValidateVoucherResponse

	if voucher.ID == 0 {
		return result, errors.New("voucher not found")
	}
	result.VoucherID = voucher.ID
	result.VoucherCode = voucher.VoucherCode

	log.Printf("transaction date %v", voucherValidateInput.FormatedTransactionDate)
	if !isVoucherActive(voucher, voucherValidateInput.FormatedTransactionDate) {
		return result, errors.New("voucher is not active")
	}
	result.VoucherStatus = "Active"

	if voucherValidateInput.TotalTransaction < voucher.MinPurchase {
		return result, errors.New("your total transaction must be greater than the minimum purchase")
	} else {
		result.TotalTransaction = voucherValidateInput.TotalTransaction
		result.TotalShippingCost = voucherValidateInput.TotalShippingCost
	}
	if voucher.PaymentMethod != "" && voucherValidateInput.PaymentMethod != voucher.PaymentMethod {
		return result, errors.New("payment method does not match")
	}

	if !isValidArea(voucher.ApplicableAreas, voucherValidateInput.Area) {
		return result, errors.New("area is not valid")
	}

	benefitAmount := CalculateBenefit(voucherValidateInput.TotalTransaction, voucherValidateInput.TotalShippingCost, voucher.DiscountAmount, voucher.VoucherCategory)
	result.BenefitAmount = benefitAmount
	return result, nil
}

func isValidArea(applicableAreasJSON, inputArea string) bool {
	var areas []string
	if err := json.Unmarshal([]byte(applicableAreasJSON), &areas); err != nil {
		return false
	}
	for _, area := range areas {
		if area == inputArea {
			return true
		}
	}
	return false
}

func isVoucherActive(voucher models.Voucher, transactionDate time.Time) bool {
	return voucher.StartDate.Before(transactionDate) && voucher.EndDate.After(transactionDate)
}
