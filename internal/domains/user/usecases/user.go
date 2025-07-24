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

func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	return u.userRepo.CreateUser(ctx, user)
}

func (u *User) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *User) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.userRepo.GetUserByEmail(ctx, email)
}
