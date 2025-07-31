package usecases

import (
	"context"

	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type User struct {
	userRepo repos.IUserRepo
}

func NewUser(userRepo repos.IUserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) CreateUser(ctx context.Context, user *model.User) error {
	return u.userRepo.CreateUser(ctx, user)
}

func (u *User) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *User) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	return u.userRepo.GetAllUsers(ctx)
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.userRepo.GetUserByEmail(ctx, email)
}

func (u *User) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return u.userRepo.GetUserByUsername(ctx, username)
}

func (u *User) AssignRole(ctx context.Context, userID string, role string) error {
	return u.userRepo.UpdateUserRole(ctx, userID, role)
}
