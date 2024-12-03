package controller

import (
	"lumoshive-be-chap41/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	// Service service.AllService
	Log    *zap.Logger
	Cacher database.Cacher
}

func NewAuthController(log *zap.Logger, rdb database.Cacher) AuthController {
	return AuthController{
		// Service: service,
		Log:    log,
		Cacher: rdb,
	}
}

func (auth *AuthController) Login(c *gin.Context) {

	// get user form database
	token := "2323232"
	IDKEY := "username-1"

	err := auth.Cacher.Set(IDKEY, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
