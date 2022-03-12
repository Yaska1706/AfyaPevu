package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaska1706/AfyaPevu/pkg/crypto"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}
