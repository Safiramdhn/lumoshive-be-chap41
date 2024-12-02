package service

import (
	"fmt"
	"lumoshive-be-chap41/models"
	"lumoshive-be-chap41/repository"
	"lumoshive-be-chap41/utils"
	"strings"
	"time"
)

type UsageService interface {
	CreateUsage(userID int, voucherInput models.VoucherDTO) error
	GetUsageVoucherByUserID(userID int) ([]models.Usage, error)
}

type usageService struct {
	UsageRepo   repository.UsageRepository
	VoucherRepo repository.VoucherRepository
}

func NewUsageService(usageRepo repository.UsageRepository, voucherRepo repository.VoucherRepository) UsageService {
	return &usageService{UsageRepo: usageRepo, VoucherRepo: voucherRepo}
}

func (s *usageService) CreateUsage(userID int, voucherInput models.VoucherDTO) error {
	voucher, err := s.VoucherRepo.GetVoucherByCode(voucherInput.VoucherCode)
	if err != nil {
		return err
	}
	const customDateFormat = "2006-01-02"
	transactionDateStr := strings.Trim(voucherInput.TransactionDate, `"`)
	parsedTransactionDate, err := time.Parse(customDateFormat, transactionDateStr)
	if err != nil {
		return fmt.Errorf("invalid transaction date format: %w", err)
	}
	voucherInput.FormatedTransactionDate = parsedTransactionDate
	voucherValidate, err := utils.ValidateVoucher(voucherInput, voucher)
	if err != nil {
		return err
	}

	usageInput := models.Usage{
		UserID:            userID,
		VoucherCode:       voucherValidate.VoucherCode,
		UsageDate:         time.Date(2024, 4, 10, 12, 0, 0, 0, time.UTC),
		BenefitAmount:     voucherValidate.BenefitAmount,
		VoucherID:         voucherValidate.VoucherID,
		TransactionAmount: voucherValidate.TotalTransaction,
	}
	return s.UsageRepo.Create(usageInput, voucher.Quantity)
}

func (s *usageService) GetUsageVoucherByUserID(userID int) ([]models.Usage, error) {
	return s.UsageRepo.GetByUserID(userID)
}
