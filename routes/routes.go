package routes

import (
	"lumoshive-be-chap41/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	// r.POST("/users", ctx.Ctl.User.CreateUser)
	// r.GET("/users/:id", ctx.Ctl.User.GetUser)
	// r.PUT("/users/:id", ctx.Ctl.User.UpdateUser)
	// r.DELETE("/users/:id", ctx.Ctl.User.DeleteUser)

	redeemRoutes(r, ctx)
	vourcherRouter(r, ctx)
	usageRouter(r, ctx)
	userRouter(r, ctx)
	return r
}

func redeemRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	redeemGroup := r.Group("/redeem")

	redeemGroup.GET("/user/:id/:voucher_id", ctx.Ctl.Redeem.RedeemVoucher)
	redeemGroup.GET("/:user_id/:voucher_type", ctx.Ctl.Redeem.GetUserRedeemByTypeVoucherController)
}

func vourcherRouter(r *gin.Engine, ctx infra.ServiceContext) {
	voucherGroup := r.Group("/voucher")

	voucherGroup.GET("/validate", ctx.Ctl.Voucher.ValidateVoucherController)
	voucherGroup.GET("/history/:voucher_code", ctx.Ctl.Voucher.GetUsageHistoryController)

	voucherGroup.POST("/", ctx.Ctl.Voucher.CreateVoucher)
	voucherGroup.DELETE("/:id", ctx.Ctl.Voucher.DeleteVoucher)
	voucherGroup.PUT("/:id", ctx.Ctl.Voucher.UpdateVoucher)
	voucherGroup.GET("/", ctx.Ctl.Voucher.GetVouchers)
	voucherGroup.GET("/point/:ratePoint", ctx.Ctl.Voucher.GetVoucherWithMinRatePoint)
}

func usageRouter(r *gin.Engine, ctx infra.ServiceContext) {
	usageGroup := r.Group("/usage")

	usageGroup.POST("/", ctx.Ctl.Usage.CreateUsageController)
	// usageGroup.GET("/:user_id", ctx.Ctl.Usage.GetUsageVoucherByUserIDController)
}

func userRouter(r *gin.Engine, ctx infra.ServiceContext) {
	userRouter := r.Group("/user")

	userRouter.GET("/redeem/:id", ctx.Ctl.User.GetUserRedeemController)
	userRouter.GET("/usage/:id", ctx.Ctl.User.GetUserUsageController)
}
