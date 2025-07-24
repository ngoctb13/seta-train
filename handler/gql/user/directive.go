package user

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ngoctb13/seta-train/internal/auth"
)

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, errors.New("access denied")
	}
	return next(ctx)
}
