package directive

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/internal/auth"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	ginContext, ok := ctx.Value("GinContextKey").(*gin.Context)
	if !ok {
		return nil, &gqlerror.Error{
			Message:    "Internal server error: Gin context not found",
			Extensions: map[string]interface{}{"code": http.StatusInternalServerError},
		}
	}

	userId, exists := ginContext.Get(auth.ContextUserIDKey)
	if !exists || userId == "" {
		return nil, &gqlerror.Error{
			Message:    "Unauthorized",
			Extensions: map[string]interface{}{"code": http.StatusUnauthorized},
		}
	}

	// role, _ := ginContext.Get(auth.ContextUserRoleKey)

	return next(ctx)
}
