package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/auth-service/internal/domain/models"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUserRole(ctx context.Context, userID string, role string) error
}
