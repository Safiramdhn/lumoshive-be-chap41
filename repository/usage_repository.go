package repository

import (
	"fmt"
	"lumoshive-be-chap41/models"

	"gorm.io/gorm"
)

type UsageRepository struct {
	DB *gorm.DB
}

func NewUsageRepository(db *gorm.DB) *UsageRepository {
	return &UsageRepository{db}
}

func (repo *UsageRepository) Create(usageInput models.Usage, voucherQuantity int) error {
	return repo.DB.Transaction(func(tx *gorm.DB) error {
		// Create a new usage record
		if err := tx.Create(&usageInput).Error; err != nil {
			return fmt.Errorf("failed to create usage record: %w", err)
		}

		// Decrement the voucher quantity
		if err := tx.Model(&models.Voucher{}).
			Where("voucher_code = ?", usageInput.VoucherCode).
			Update("quantity", gorm.Expr("quantity - ?", 1)).Error; err != nil {
			return fmt.Errorf("failed to update voucher quantity: %w", err)
		}

		return nil
	})
}

func (repo *UsageRepository) GetByUserID(userID int) ([]models.Usage, error) {
	var usages []models.Usage

	err := repo.DB.
		Preload("Voucher").
		Preload("User").
		Where("user_id = ?", userID).
		Find(&usages).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get usages by user ID: %w", err)
	}

	return usages, nil
}

func (repo *UsageRepository) GetByVoucherCode(voucherCode string) (*models.Voucher, error) {
	var voucher models.Voucher

	if err := repo.DB.Preload("Usage").Where("voucher_code =?", voucherCode).Find(&voucher).Error; err != nil {
		return nil, fmt.Errorf("failed to get usages by user ID: %w", err)
	}

	return &voucher, nil
}
