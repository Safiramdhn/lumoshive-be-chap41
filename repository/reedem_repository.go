package repository

import (
	"errors"
	"lumoshive-be-chap41/models"
	"time"

	"gorm.io/gorm"
)

type RedeemRepository struct {
	DB *gorm.DB
}

func NewRedeemRepository(db *gorm.DB) *RedeemRepository {
	return &RedeemRepository{db}
}

func (repo *RedeemRepository) GetUserRedeem(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	var redeems []models.Redeem

	err := repo.DB.Preload("Voucher", "voucher_type = ?", voucherFilter.VoucherType).
		Where("user_id = ?", userID).
		Find(&redeems).Error
	if err != nil {
		return nil, err
	}

	now := time.Now()
	activeRedeems := []models.Redeem{}
	for _, redeem := range redeems {
		if redeem.Voucher.ID != 0 && redeem.Voucher.StartDate.Before(now) && redeem.Voucher.EndDate.After(now) {
			activeRedeems = append(activeRedeems, redeem)
		}
	}

	return activeRedeems, nil
}

func (repo *RedeemRepository) GetUserRedeemByType(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	var redeems []models.Redeem

	err := repo.DB.Preload("Voucher", "voucher_type = ?", voucherFilter.VoucherType).
		Where("user_id = ?", userID).
		Find(&redeems).Error
	if err != nil {
		return nil, err
	}
	return redeems, nil
}

func (repo *RedeemRepository) GetAllUserRedeems(userID int) ([]models.Redeem, error) {
	var redeems []models.Redeem
	err := repo.DB.Where("user_id =?", userID).Find(&redeems).Error
	if err != nil {
		return nil, err
	}
	return redeems, nil
}

func (repo *RedeemRepository) RedeemVoucher(user *models.User, voucherID int) (models.Redeem, error) {
	var voucher models.Voucher
	if err := repo.DB.First(&voucher, voucherID).Error; err != nil {
		return models.Redeem{}, errors.New("voucher not found")
	}

	activeRedeems, err := repo.GetUserRedeem(user.ID, voucher)
	if err != nil {
		return models.Redeem{}, err
	}

	for _, voucher := range activeRedeems {
		if voucher.VoucherID == voucherID {
			return models.Redeem{}, errors.New("only one code in voucher for customers used")
		}
	}

	if voucher.MinRatePoint > user.Points {
		return models.Redeem{}, errors.New("not enough points to redeem")
	}

	redeem := models.Redeem{
		UserID:      user.ID,
		VoucherID:   voucherID,
		VoucherCode: voucher.VoucherCode,
		RedeemDate:  time.Now(),
		Voucher:     voucher,
	}

	if err := repo.DB.Create(&redeem).Error; err != nil {
		return models.Redeem{}, err
	}

	user.Points = user.Points - voucher.MinRatePoint
	return redeem, nil
}
