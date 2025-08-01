package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/internal/auth/grahpql"
)

const (
	userIDKey   = "userID"
	userRoleKey = "userRole"
)

var authService = grahpql.NewAuthService()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		userInfo, err := authService.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set(userIDKey, userInfo.ID)
		c.Set(userRoleKey, userInfo.Role)
		c.Next()
	}
}
