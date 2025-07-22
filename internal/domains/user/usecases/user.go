package usecases

import (
	"context"

	"github.com/ngoctb13/seta-train/config"
	"github.com/ngoctb13/seta-train/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/internal/models"
)

type User struct {
	cfg      *config.AppConfig
	userRepo repos.IUserRepo
}

func NewUser(cfg *config.AppConfig, userRepo repos.IUserRepo) *User {
	return &User{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *User) CreateUser(ctx context.Context, user *models.User) error {
	return u.userRepo.CreateUser(ctx, user)
}
