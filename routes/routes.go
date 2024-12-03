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

	r.GET("/login", ctx.Ctl.Auth.Login)
	redeemRoutes(r, ctx)
	voucherRouter(r, ctx)
	usageRouter(r, ctx)
	userRouter(r, ctx)
	return r
}

func redeemRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	redeemGroup := r.Group("/redeem")

	redeemGroup.GET("/user/:id/:voucher_id", ctx.Ctl.Redeem.RedeemVoucher)
	redeemGroup.GET("/:user_id/:voucher_type", ctx.Ctl.Redeem.GetUserRedeemByTypeVoucherController)
}

func voucherRouter(router *gin.Engine, ctx infra.ServiceContext) {
	// Define the voucher group and apply authentication middleware
	voucherGroup := router.Group("/voucher", ctx.Middleware.Authentication())

	// Validation and usage history routes
	voucherGroup.GET("/validate", ctx.Ctl.Voucher.ValidateVoucherController)
	voucherGroup.GET("/history/:voucher_code", ctx.Ctl.Voucher.GetUsageHistoryController)

	// CRUD operations for vouchers
	voucherGroup.POST("/", ctx.Ctl.Voucher.CreateVoucher)      // Create a voucher
	voucherGroup.GET("/", ctx.Ctl.Voucher.GetVouchers)         // List all vouchers
	voucherGroup.PUT("/:id", ctx.Ctl.Voucher.UpdateVoucher)    // Update a voucher by ID
	voucherGroup.DELETE("/:id", ctx.Ctl.Voucher.DeleteVoucher) // Delete a voucher by ID

	// Additional routes
	voucherGroup.GET("/point/:ratePoint", ctx.Ctl.Voucher.GetVoucherWithMinRatePoint) // Get vouchers by minimum rate points
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
	userRouter.GET("/validate-voucher", ctx.Ctl.Voucher.ValidateVoucherController)

}
