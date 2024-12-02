package controller

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"lumoshive-be-chap41/models"
// 	"lumoshive-be-chap41/service"
// 	"lumoshive-be-chap41/utils"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/zap"
// )

// var response utils.HTTPResponse

// // Utility function to create a controller instance
// func setupTestController() (*UserController, *service.MockUserService, *zap.Logger) {
// 	mockService := new(service.MockUserService)
// 	logger, _ := zap.NewProduction()
// 	return NewUserController(mockService, logger), mockService, logger
// }

// // Test for successful retrieval of user
// func TestGetUserSuccess(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	controller, mockService, logger := setupTestController()
// 	defer logger.Sync()

// 	mockUser := models.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
// 	mockService.On("GetUser", uint(1)).Return(mockUser, nil)

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
// 	c.Request, _ = http.NewRequest("GET", "/users/1", nil)

// 	controller.GetUser(c)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	// var response HTTPResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.True(t, response.Success)
// 	assert.Equal(t, "User retrieved successfully", response.Description)

// 	// Casting response.Data to map to access the fields
// 	// dataMap, ok := response.Data.(map[string]interface{})
// 	if assert.True(t, ok) {
// 		assert.Equal(t, mockUser.Name, dataMap["Name"])
// 	}
// }

// // Test for user not found in retrieval
// func TestGetUserNotFound(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	controller, mockService, logger := setupTestController()
// 	defer logger.Sync()

// 	mockService.On("GetUser", uint(1)).Return(models.User{}, assert.AnError)

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
// 	c.Request, _ = http.NewRequest("GET", "/users/1", nil)

// 	controller.GetUser(c)

// 	assert.Equal(t, http.StatusNotFound, w.Code)
// 	// var response HTTPResponse
// 	// json.Unmarshal(w.Body.Bytes(), &response)
// 	// assert.False(t, response.Success)
// 	// assert.Equal(t, "NOT_FOUND", response.ErrorCode)
// }

// // Test for invalid ID in user retrieval
// func TestGetUserInvalidID(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	controller, _, logger := setupTestController()
// 	defer logger.Sync()

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc"}}
// 	c.Request, _ = http.NewRequest("GET", "/users/abc", nil)

// 	// controller.GetUser(c)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	// var response HTTPResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.False(t, response.Success)
// 	assert.Equal(t, "INVALID_ID", response.ErrorCode)
// }
