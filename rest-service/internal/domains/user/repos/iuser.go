package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]*model.User, error)
}
