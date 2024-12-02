package repository

import "gorm.io/gorm"

type Repository struct {
	User    UserRepository
	Voucher VoucherRepository
	Redeem  RedeemRepository
	Usage   UsageRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		User:    *NewUserRepository(db),
		Voucher: *NewVoucherRepository(db),
		Redeem:  *NewRedeemRepository(db),
		Usage:   *NewUsageRepository(db),
	}
}
