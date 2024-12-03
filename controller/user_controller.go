package controller

import (
	"lumoshive-be-chap41/service"
	"lumoshive-be-chap41/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserController(service service.UserService, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

// GetUserRedeemController godoc
// @Summary Retrieve a list of redeems for a specific user
// @Description Retrieves the redeem data for a given user by their ID.
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Redeem "List of user redeems and success message"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/redeem/{id} [get]

func (ctrl *UserController) GetUserRedeemController(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}
	userRedeems, err := ctrl.service.GetUserRedeem(userID)
	if err != nil {
		ctrl.logger.Error("Failed to get user redeems", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", "Failed to get user redeems", http.StatusInternalServerError)
		return
	}
	ctrl.logger.Info("User redeem retrieved successfully", zap.Any("user_redeems", userRedeems))
	utils.ResponseOK(c, userRedeems, "user redeem retrieved successfully")
}

// GetUserUsageController godoc
// @Summary Retrieve a list of usage data for a specific user
// @Description Retrieves the usage data for a given user by their ID.
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Usage "List of user usage and success message"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/usage/{id} [get]
func (ctrl *UserController) GetUserUsageController(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusBadRequest)
		return
	}
	userUsage, err := ctrl.service.GetUserUsage(userID)
	if err != nil {
		ctrl.logger.Error("Failed to get user usage", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", "Failed to get user usage", http.StatusInternalServerError)
		return
	}
	ctrl.logger.Info("User usage retrieved successfully", zap.Any("user_usage", userUsage))
	utils.ResponseOK(c, userUsage, "user usage retrieved successfully")
}

// func (ctrl *UserController) CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
// 		// responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := ctrl.service.CreateUser(user); err != nil {
// 		ctrl.logger.Error("Failed to create user", zap.Error(err))
// 		// responseError(c, "CREATE_ERROR", "Failed to create user", http.StatusInternalServerError)
// 		return
// 	}

// 	ctrl.logger.Info("User created successfully", zap.Any("user", user))
// 	// responseOK(c, user, "User created successfully")
// }

// func (ctrl *UserController) GetUser(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		ctrl.logger.Error("Invalid user ID", zap.Error(err))
// 		// responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := ctrl.service.GetUser(id)
// 	if err != nil {
// 		ctrl.logger.Error("User not found", zap.Error(err))
// 		// responseError(c, "NOT_FOUND", "User not found", http.StatusNotFound)
// 		return
// 	}

// 	ctrl.logger.Info("User retrieved successfully", zap.Any("user", user))
// 	// responseOK(c, user, "User retrieved successfully")
// }

// func (ctrl *UserController) UpdateUser(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		ctrl.logger.Error("Invalid user ID", zap.Error(err))
// 		// responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		ctrl.logger.Error("Failed to bind user data", zap.Error(err))
// 		// responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	user.ID = id

// 	if err := ctrl.service.UpdateUser(user); err != nil {
// 		ctrl.logger.Error("Failed to update user", zap.Error(err))
// 		// responseError(c, "UPDATE_ERROR", "Failed to update user", http.StatusInternalServerError)
// 		return
// 	}

// 	ctrl.logger.Info("User updated successfully", zap.Any("user", user))
// 	// responseOK(c, user, "User updated successfully")
// }

// func (ctrl *UserController) DeleteUser(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		ctrl.logger.Error("Invalid user ID", zap.Error(err))
// 		// responseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
// 		return
// 	}

// 	if err := ctrl.service.DeleteUser(uint(id)); err != nil {
// 		ctrl.logger.Error("Failed to delete user", zap.Error(err))
// 		// responseError(c, "DELETE_ERROR", "Failed to delete user", http.StatusInternalServerError)
// 		return
// 	}

// 	ctrl.logger.Info("User deleted successfully", zap.Int("userID", id))
// 	// responseOK(c, nil, "User deleted successfully")
// }
