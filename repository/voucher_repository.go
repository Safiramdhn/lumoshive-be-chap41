package repository

import (
	"errors"
	"fmt"
	"lumoshive-be-chap41/models"
	"time"

	"gorm.io/gorm"
)

type VoucherRepositoryIn interface {
	GetUserVoucher(voucherFilter models.Voucher) (models.Voucher, error)
	GetVoucherByCode(voucherCode string) (models.Voucher, error)
	CreateVoucher(voucher *models.Voucher) error
	DeleteVoucher(id int) error
	UpdateVoucher(voucher *models.Voucher) error
	GetVouchers(filter map[string]interface{}) ([]models.Voucher, error)
	GetVoucherWithMinRatePoint(ratePoint int) ([]models.Voucher, error)
}

type VoucherRepository struct {
	DB *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{db}
}

func (repo *VoucherRepository) GetUserVoucher(voucherFilter models.Voucher) (models.Voucher, error) {
	var voucher models.Voucher
	err := repo.DB.Where("voucher_code = ? AND voucher_type = ?", voucherFilter.VoucherCode, voucherFilter.VoucherType).First(&voucher).Error
	return voucher, err
}

func (repo *VoucherRepository) GetVoucherByCode(voucherCode string) (models.Voucher, error) {
	var voucher models.Voucher
	err := repo.DB.Where("voucher_code = ?", voucherCode).First(&voucher).Error
	if err == gorm.ErrRecordNotFound {
		return voucher, nil
	}
	return voucher, err
}

func (r *VoucherRepository) CreateVoucher(voucher *models.Voucher) error {
	return r.DB.Create(voucher).Error
}

func (r *VoucherRepository) DeleteVoucher(id int) error {
	if result := r.DB.Delete(&models.Voucher{}, id); result.RowsAffected == 0 {
		return errors.New("voucher not found")
	}
	return nil
}

func (r *VoucherRepository) UpdateVoucher(voucher *models.Voucher) error {
	return r.DB.Save(voucher).Error
}

func (r *VoucherRepository) GetVouchers(filter map[string]interface{}) ([]models.Voucher, error) {
	var vouchers []models.Voucher
	query := r.DB.Model(&models.Voucher{})

	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}

	err := query.Find(&vouchers).Error
	return vouchers, err
}

func (r *VoucherRepository) GetVoucherWithMinRatePoint(ratePoint int) ([]map[string]interface{}, error) {
	var vouchers []map[string]interface{}
	result := r.DB.Model(&models.Voucher{}).
		Select("voucher_name, discount_amount, min_rate_point").
		Where("min_rate_point <= ? AND end_date >= ?", ratePoint, time.Now()).
		Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return vouchers, nil
}

func (repo *VoucherRepository) GetUserUsageByVoucherCode(voucherCode string) (*models.Voucher, error) {
	if voucherCode == "" {
		return nil, fmt.Errorf("voucher code cannot be empty")
	}

	var voucher models.Voucher

	err := repo.DB.
		Preload("Usages.User"). // Preload nested relations in one step
		Where("voucher_code = ?", voucherCode).
		First(&voucher).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no record is found without an error
		}
		return nil, fmt.Errorf("failed to fetch voucher by code %s: %w", voucherCode, err)
	}

	return &voucher, nil
}
