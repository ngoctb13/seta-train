package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ngoctb13/seta-train/auth-service/graph/model"
	"github.com/ngoctb13/seta-train/auth-service/internal/auth"
	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/shared-modules/logger"
)

type contextKey string

const (
	userIDKey   contextKey = "userID"
	userRoleKey contextKey = "userRole"
)

type Resolver struct {
	UserUsecase *usecases.User
	Logger      *logger.Logger
}

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role *model.Role) (interface{}, error) {
	header, ok := ctx.Value("Authorization").(string)
	if !ok || header == "" {
		return nil, errors.New("missing Authorization header")
	}
	userID, userRole, err := auth.ParseAndValidateJWT(header)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}
	if role != nil && userRole != string(*role) {
		return nil, errors.New("forbidden: insufficient role")
	}
	ctx = context.WithValue(ctx, userIDKey, userID)
	ctx = context.WithValue(ctx, userRoleKey, userRole)
	return next(ctx)
}
