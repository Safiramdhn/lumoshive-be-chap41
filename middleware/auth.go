package middleware

import (
	"lumoshive-be-chap41/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Cacher database.Cacher
}

func NewMiddleware(cacher database.Cacher) Middleware {
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("ID-KEY")
		val, err := m.Cacher.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if val == "" || val != token {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		// before request
		c.Next()

	}
}
