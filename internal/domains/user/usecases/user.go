package usecases

import (
	"context"

	"github.com/ngoctb13/seta-train/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/internal/models"
)

type User struct {
	userRepo repos.IUserRepo
}

func NewUser(userRepo repos.IUserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}

// Todo: if have error, should log before return error
func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	return u.userRepo.CreateUser(ctx, user)
}

func (u *User) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *User) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}
