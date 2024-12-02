package controller

import (
	"lumoshive-be-chap41/models"
	"lumoshive-be-chap41/service"
	"lumoshive-be-chap41/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedeemController struct {
	service service.RedeemService
	user    service.UserService
	logger  *zap.Logger
}

func NewRedeemController(service service.RedeemService, user service.UserService, logger *zap.Logger) *RedeemController {
	return &RedeemController{service, user, logger}
}

func (ctrl *RedeemController) GetUserRedeemByTypeVoucherController(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": "INVALID_ID", "error_message": "Invalid user ID"})
		return
	}

	voucherType := c.Param("voucher_type")
	if voucherType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error_code": "EMPTY_PARAM", "error_message": "voucher type is empty"})
		return
	}

	voucherFilter := models.Voucher{VoucherType: voucherType}
	redeems, err := ctrl.service.GetActiveUserRedeems(userID, voucherFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code":    "INTERNAL_SERVER_ERROR",
			"error_message": err.Error(),
		})
		return
	}

	if len(redeems) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":        []models.Redeem{},
			"description": "user has no redeem voucher",
			"success":     true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    redeems,
		"message": "user redeem successfully retrieved",
	})
}

func (ctrl *RedeemController) RedeemVoucher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	VoucherId, err := strconv.Atoi(c.Param("voucher_id"))
	if err != nil {
		ctrl.logger.Error("Invalid Voucher Id", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctrl.user.GetUser(id)
	if err != nil {
		ctrl.logger.Error("User not found", zap.Error(err))
		utils.ResponseError(c, "NOT_FOUND", "User not found", http.StatusNotFound)
		return
	}

	reedem, err := ctrl.service.RedeemVoucher(&user, VoucherId)
	if err != nil {
		ctrl.logger.Error("Reedem voucher error", zap.Error(err))
		utils.ResponseError(c, "REEDEM_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	err = ctrl.user.UpdateUser(user)
	if err != nil {
		ctrl.logger.Error("Error update point user", zap.Error(err))
		utils.ResponseError(c, "ERR0R_UPDATE_POINT_USER", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Reedem voucher successfully")
	utils.ResponseOK(c, reedem, "Reedem voucher successfully")
}
