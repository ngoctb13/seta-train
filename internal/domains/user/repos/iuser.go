package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/internal/models"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
