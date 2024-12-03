package middleware

import (
	"log"
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

		// Log the received headers for debugging
		log.Printf("Authentication Middleware: Received token: %s, ID-KEY: %s", token, id)

		// Check if ID-KEY header is present
		if id == "" {
			log.Println("Authentication Middleware: Missing ID-KEY header")
			c.JSON(http.StatusBadRequest, "Missing ID-KEY header")
			c.Abort()
			return
		}

		// Fetch value from the cacher
		val, err := m.Cacher.Get(id)
		if err != nil {
			log.Printf("Authentication Middleware: Error retrieving cache for ID-KEY %s: %v", id, err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			c.Abort()
			return
		}

		log.Printf("Authentication Middleware: Cache value for ID-KEY %s: %s", id, val)

		// Validate token
		if val == "" || val != token {
			log.Printf("Authentication Middleware: Token mismatch or missing for ID-KEY %s. Expected: %s, Received: %s", id, val, token)
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		// Before proceeding to the next handler
		log.Println("Authentication Middleware: Request authorized")
		c.Next()
	}
}
