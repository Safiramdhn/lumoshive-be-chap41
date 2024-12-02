package service

import (
	"lumoshive-be-chap41/models"
	"lumoshive-be-chap41/repository"
	"time"
)

type RedeemService interface {
	GetActiveUserRedeems(userID int, voucherFilter models.Voucher) ([]models.Redeem, error)
	GetAllUserRedeems(userID int) ([]models.Redeem, error)
	RedeemVoucher(user *models.User, voucherID int) (models.Redeem, error)
}

type redeemService struct {
	Repo repository.RedeemRepository
}

func NewRedeemService(repo repository.RedeemRepository) RedeemService {
	return &redeemService{repo}
}

func (s *redeemService) GetActiveUserRedeems(userID int, voucherFilter models.Voucher) ([]models.Redeem, error) {
	redeems, err := s.Repo.GetUserRedeemByType(userID, voucherFilter)
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

func (s *redeemService) GetAllUserRedeems(userID int) ([]models.Redeem, error) {
	redeems, err := s.Repo.GetAllUserRedeems(userID)
	if err != nil {
		return nil, err
	}
	return redeems, nil
}

func (s *redeemService) RedeemVoucher(user *models.User, voucherID int) (models.Redeem, error) {
	return s.Repo.RedeemVoucher(user, voucherID)
}
