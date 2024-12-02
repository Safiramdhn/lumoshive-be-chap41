package controller

import (
	"lumoshive-be-chap41/service"

	"go.uber.org/zap"
)

type Controller struct {
	User    UserController
	Redeem  RedeemController
	Voucher VoucherController
	Usage   UsageController
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		User:    *NewUserController(service.User, logger),
		Redeem:  *NewRedeemController(service.Reedem, service.User, logger),
		Voucher: *NewVoucherController(service.Voucher, logger),
		Usage:   *NewUsageController(service.Usage, logger),
	}
}
