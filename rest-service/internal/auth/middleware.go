package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userIDKey   = "userID"
	userRoleKey = "userRole"
)

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func verifyTokenWithAuthService(token string) (*UserInfo, error) {
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "http://localhost:8080/query"
	}

	gqlQuery := map[string]interface{}{
		"query":     "query VerifyToken($token: String!) { verifyToken(token: $token) { id username email password role } }",
		"variables": map[string]interface{}{"token": token},
	}
	jsonBody, _ := json.Marshal(gqlQuery)
	resp, err := http.Post(authServiceURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Parse response
	var gqlResp struct {
		Data struct {
			VerifyToken *UserInfo `json:"verifyToken"`
		} `json:"data"`
		Errors []interface{} `json:"errors"`
	}
	if err := json.Unmarshal(body, &gqlResp); err != nil {
		return nil, err
	}
	if len(gqlResp.Errors) > 0 || gqlResp.Data.VerifyToken == nil {
		return nil, err
	}
	return gqlResp.Data.VerifyToken, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		userInfo, err := verifyTokenWithAuthService(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set(userIDKey, userInfo.ID)
		c.Set(userRoleKey, userInfo.Role)
		c.Next()
	}
}
