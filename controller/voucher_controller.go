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

type VoucherController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewVoucherController(service service.VoucherService, logger *zap.Logger) *VoucherController {
	return &VoucherController{service, logger}
}

// ValidateVoucherController godoc
// @Summary Validate a voucher
// @Description Validate the voucher with the provided voucher data
// @Tags Voucher
// @Accept json
// @Produce json
// @Param voucher body models.VoucherDTO true "Voucher DTO"
// @Success 200 {object} models.VoucherDTO
// @Failure 400 {object} utils.HTTPResponse
// @Router /voucher/validate [get]
func (ctrl *VoucherController) ValidateVoucherController(c *gin.Context) {
	var voucherInput models.VoucherDTO
	if err := c.ShouldBindJSON(&voucherInput); err != nil {
		ctrl.logger.Error("Failed to bind voucher data", zap.Error(err))
		utils.ResponseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	validateResult, err := ctrl.service.ValidateVoucher(voucherInput)
	if err != nil {
		ctrl.logger.Error("Invalid voucher data", zap.Error(err))
		utils.ResponseError(c, "INVALID_DATA", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.logger.Info("Voucher data validated successfully")
	utils.ResponseOK(c, validateResult, "Voucher data validated successfully")
}

// @Summary Create Voucher
// @Description Create a new voucher
// @Tags Voucher
// @Accept json
// @Produce json
// @Param voucher body models.Voucher true "Voucher Data"
// @Success 200 {object} utils.HTTPResponse "Voucher created successfully"
// @Failure 400 {object} utils.HTTPResponse "Invalid input data"
// @Failure 500 {object} utils.HTTPResponse "Internal server error"
// @Router /voucher [post]
func (ctrl *VoucherController) CreateVoucher(c *gin.Context) {
	var voucher models.Voucher
	var err error

	if err = c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to create voucher", zap.Error(err))
		utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	// voucher.StartDate, err = utils.TimeDateParse(voucher.StartDate.String())
	// if err != nil {
	// 	ctrl.logger.Error("Failed to create voucher", zap.Error(err))
	// 	utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// voucher.EndDate, err = utils.TimeDateParse(voucher.EndDate.String())
	// if err != nil {
	// 	ctrl.logger.Error("Failed to create voucher", zap.Error(err))
	// 	utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
	// 	return
	// }

	if err := ctrl.service.CreateVoucher(&voucher); err != nil {
		ctrl.logger.Error("Failed to create voucher", zap.Error(err))
		utils.ResponseError(c, "CREATE_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Create voucher successfully")
	utils.ResponseOK(c, voucher, "Create voucher successfully")
}

// @Summary Delete Voucher
// @Description Delete a voucher by ID
// @Tags Voucher
// @Param id path int true "Voucher ID"
// @Success 200 {object} utils.HTTPResponse "Voucher deleted successfully"
// @Failure 400 {object} utils.HTTPResponse "Invalid ID"
// @Failure 404 {object} utils.HTTPResponse "Voucher not found"
// @Router /voucher/{id} [delete]
func (ctrl *VoucherController) DeleteVoucher(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Failed to delete voucher", zap.Error(err))
		utils.ResponseError(c, "DELETE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.DeleteVoucher(id); err != nil {
		ctrl.logger.Error("Data not found", zap.Error(err))
		utils.ResponseError(c, "DATA_NOT_FOUND", err.Error(), http.StatusNotFound)
		return
	}

	ctrl.logger.Info("Voucher deleted successfully")
	utils.ResponseOK(c, nil, "Voucher deleted successfully")
}

// @Summary Update Voucher
// @Description Update an existing voucher
// @Tags Voucher
// @Accept json
// @Produce json
// @Param voucher body models.Voucher true "Voucher Data"
// @Success 200 {object} utils.HTTPResponse "Voucher updated successfully"
// @Failure 400 {object} utils.HTTPResponse "Invalid input data"
// @Failure 500 {object} utils.HTTPResponse "Internal server error"
// @Router /voucher [put]
func (ctrl *VoucherController) UpdateVoucher(c *gin.Context) {
	var voucher models.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		ctrl.logger.Error("Failed to update voucher", zap.Error(err))
		utils.ResponseError(c, "UPDATE_VOUCHER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctrl.service.UpdateVoucher(&voucher); err != nil {
		ctrl.logger.Error("Failed to update voucher", zap.Error(err))
		utils.ResponseError(c, "UPDATE_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Voucher update successfully")
	utils.ResponseOK(c, voucher, "Voucher update successfully")
}

// @Summary Get Vouchers
// @Description Retrieve vouchers with optional filters
// @Tags Voucher
// @Produce json
// @Param voucher_code query string false "Filter by voucher code"
// @Param voucher_type query string false "Filter by voucher type"
// @Success 200 {array} models.VoucherWithStatus "List of vouchers"
// @Failure 500 {object} utils.HTTPResponse "Internal server error"
// @Router /vouchers [get]
func (ctrl *VoucherController) GetVouchers(c *gin.Context) {
	filter := make(map[string]interface{})
	if c.Query("voucher_code") != "" {
		filter["voucher_code"] = c.Query("voucher_code")
	}
	if c.Query("voucher_type") != "" {
		filter["voucher_type"] = c.Query("voucher_type")
	}

	vouchers, err := ctrl.service.GetVouchers(filter)
	if err != nil {
		ctrl.logger.Error("Failed to get voucher", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	var response []models.VoucherWithStatus
	for _, voucher := range vouchers {
		response = append(response, models.VoucherWithStatus{
			Voucher:  voucher,
			IsActive: voucher.IsActive(),
		})
	}

	ctrl.logger.Info("Get voucher successfully")
	utils.ResponseOK(c, response, "Get voucher successfully")
}

// @Summary Get Voucher by Minimum Rate Point
// @Description Retrieve vouchers with a minimum rate point
// @Tags Voucher
// @Param ratePoint path int true "Minimum rate point"
// @Success 200 {array} models.Voucher "List of vouchers"
// @Failure 400 {object} utils.HTTPResponse "Invalid rate point"
// @Failure 404 {object} utils.HTTPResponse "No vouchers found"
// @Failure 500 {object} utils.HTTPResponse "Internal server error"
// @Router /vouchers/min-rate/{ratePoint} [get]
func (ctrl *VoucherController) GetVoucherWithMinRatePoint(c *gin.Context) {
	ratePoint, err := strconv.Atoi(c.Param("ratePoint"))
	if err != nil {
		ctrl.logger.Error("Failed to parse ratePoint", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", "Invalid ratePoint parameter", http.StatusBadRequest)
		return
	}

	vouchers, err := ctrl.service.GetVoucherWithMinRatePoint(ratePoint)
	if err != nil {
		ctrl.logger.Error("Failed to get vouchers", zap.Error(err))
		utils.ResponseError(c, "GET_VOUCHER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	if len(vouchers) == 0 {
		ctrl.logger.Error("No vouchers found")
		utils.ResponseError(c, "GET_VOUCHER_ERROR", "No vouchers found", http.StatusNotFound)
		return
	}

	ctrl.logger.Info("Get vouchers successfully")
	utils.ResponseOK(c, vouchers, "Get vouchers successfully")
}

// @Summary Get Voucher Usage History
// @Description Retrieve usage history of a voucher by its code
// @Tags Voucher
// @Param voucher_code path string true "Voucher Code"
// @Success 200 {object} utils.HTTPResponse "Voucher usage history retrieved"
// @Failure 400 {object} utils.HTTPResponse "Empty voucher code"
// @Failure 500 {object} utils.HTTPResponse "Internal server error"
// @Router /voucher/usage/{voucher_code} [get]
func (ctrl *VoucherController) GetUsageHistoryController(c *gin.Context) {
	voucherCode := c.Param("voucher_code")
	if voucherCode == "" {
		ctrl.logger.Error("voucher code is empty")
		utils.ResponseError(c, "EMPTY_PARAM", "voucher code is empty", http.StatusBadRequest)
		return
	}

	voucher, err := ctrl.service.GetVoucherUsageHistory(voucherCode)
	if err != nil {
		ctrl.logger.Error("Failed to get voucher usage history", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.logger.Info("Voucher usage history retrieved successfully")
	utils.ResponseOK(c, voucher, "voucher usage history retrieved successfully")
}
