package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPResponse defines the structure for standard API responses
type HTTPResponse struct {
	Success     bool        `json:"success" example:"true"`
	ErrorCode   string      `json:"error_code,omitempty" example:"ERR123"`
	Description string      `json:"description,omitempty" example:"Operation successful"`
	Data        interface{} `json:"data,omitempty" swaggertype:"object"` // Use swaggertype for unsupported types
}

func ResponseOK(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, HTTPResponse{
		Success:     true,
		Description: description,
		Data:        data,
	})
}

func ResponseError(c *gin.Context, errorCode string, description string, httpStatusCode int) {
	c.JSON(httpStatusCode, HTTPResponse{
		Success:     false,
		ErrorCode:   errorCode,
		Description: description,
	})
}
