package grahpql

import (
	"fmt"
)

const (
	verifyTokenQuery = `
		query VerifyToken($token: String!) { 
			verifyToken(token: $token) { 
				id 
				username 
				email 
				password 
				role 
			} 
		}
	`
)

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthService struct {
	client *GraphQLClient
}

func NewAuthService() *AuthService {
	return &AuthService{
		client: NewGraphQLClient(),
	}
}

// VerifyTokenResponse represents the response structure for verifyToken query
type VerifyTokenResponse struct {
	VerifyToken *UserInfo `json:"verifyToken"`
}

// VerifyToken verifies a JWT token and returns user information
func (a *AuthService) VerifyToken(token string) (*UserInfo, error) {
	var response VerifyTokenResponse

	err := a.client.ExecuteQuery(verifyTokenQuery, map[string]interface{}{
		"token": token,
	}, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}

	if response.VerifyToken == nil {
		return nil, fmt.Errorf("invalid token")
	}

	return response.VerifyToken, nil
}
