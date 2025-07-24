package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/internal/models"
)

type contextKey string

const userIDKey = contextKey("userID")
const userRoleKey = contextKey("userRole")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
			return
		}
		role, _ := claims["role"].(string)
		c.Set(string(userIDKey), userId)
		c.Set(string(userRoleKey), role)
		c.Next()
	}
}

func WithUser(ctx context.Context, userID string, role string) context.Context {
	ctx = context.WithValue(ctx, userIDKey, userID)
	ctx = context.WithValue(ctx, userRoleKey, role)
	return ctx
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

func GetUserRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(userRoleKey).(string)
	return role, ok
}

func ForContext(ctx context.Context) *models.User {
	userID, ok := GetUserID(ctx)
	if !ok || userID == "" {
		return nil
	}
	role, _ := GetUserRole(ctx)
	return &models.User{
		UserID: userID,
		Role:   role,
	}
}
