package service

import (
	"lumoshive-be-chap41/repository"
)

type Service struct {
	User    UserService
	Reedem  RedeemService
	Voucher VoucherService
	Usage   UsageService
}

func NewService(repo repository.Repository) Service {
	return Service{
		User:    NewUserService(repo.User),
		Reedem:  NewRedeemService(repo.Redeem),
		Voucher: NewVoucherService(repo.Voucher),
		Usage:   NewUsageService(repo.Usage, repo.Voucher),
	}
}
