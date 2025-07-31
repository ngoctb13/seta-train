package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
	UpdateUserRole(ctx context.Context, userID string, role string) error
}
