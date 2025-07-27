package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ngoctb13/seta-train/auth-service/graph/model"
	"github.com/ngoctb13/seta-train/auth-service/internal/auth"
	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUsecase *usecases.User
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
	if role != nil && string(*role) != userRole {
		return nil, errors.New("forbidden: insufficient role")
	}
	ctx = context.WithValue(ctx, "userID", userID)
	ctx = context.WithValue(ctx, "userRole", userRole)
	return next(ctx)
}
