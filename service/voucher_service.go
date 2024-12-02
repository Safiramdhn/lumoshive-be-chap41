package service

import (
	"fmt"
	"lumoshive-be-chap41/models"
	"lumoshive-be-chap41/repository"
	"lumoshive-be-chap41/utils"
	"strings"
	"time"
)

type VoucherService interface {
	ValidateVoucher(voucherInput models.VoucherDTO) (*models.ValidateVoucherResponse, error)
	CreateVoucher(voucher *models.Voucher) error
	DeleteVoucher(id int) error
	UpdateVoucher(voucher *models.Voucher) error
	GetVouchers(filter map[string]interface{}) ([]models.Voucher, error)
	GetVoucherWithMinRatePoint(ratePoint int) ([]map[string]interface{}, error)
	GetVoucherUsageHistory(voucherCode string) (*models.Voucher, error)
}

type voucherService struct {
	Repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo}
}

func (s *voucherService) ValidateVoucher(voucherInput models.VoucherDTO) (*models.ValidateVoucherResponse, error) {
	voucher, err := s.Repo.GetVoucherByCode(voucherInput.VoucherCode)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve voucher: %w", err)
	}

	// Parse transaction date
	const customDateFormat = "2006-01-02"
	transactionDateStr := strings.Trim(voucherInput.TransactionDate, `"`)
	parsedTransactionDate, err := time.Parse(customDateFormat, transactionDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction date format: %w", err)
	}
	voucherInput.FormatedTransactionDate = parsedTransactionDate

	// Validate the voucher
	validationResult, err := utils.ValidateVoucher(voucherInput, voucher)
	if err != nil {
		return nil, fmt.Errorf("voucher validation failed: %w", err)
	}

	return &validationResult, nil
}

func (s *voucherService) GetVoucherUsageHistory(voucherCode string) (*models.Voucher, error) {
	return s.Repo.GetUserUsageByVoucherCode(voucherCode)
}

func (s *voucherService) CreateVoucher(voucher *models.Voucher) error {
	return s.Repo.CreateVoucher(voucher)
}

func (s *voucherService) DeleteVoucher(id int) error {
	return s.Repo.DeleteVoucher(id)
}

func (s *voucherService) UpdateVoucher(voucher *models.Voucher) error {
	return s.Repo.UpdateVoucher(voucher)
}

func (s *voucherService) GetVouchers(filter map[string]interface{}) ([]models.Voucher, error) {
	return s.Repo.GetVouchers(filter)
}

func (s *voucherService) GetVoucherWithMinRatePoint(ratePoint int) ([]map[string]interface{}, error) {
	return s.Repo.GetVoucherWithMinRatePoint(ratePoint)
}
