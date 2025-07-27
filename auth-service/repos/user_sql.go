package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
	"gorm.io/gorm"
)

type userSQLRepo struct {
	db *gorm.DB
}

func NewUserSQLRepo(db *gorm.DB) *userSQLRepo {
	return &userSQLRepo{
		db: db,
	}
}

func (u *userSQLRepo) CreateUser(ctx context.Context, user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userSQLRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, "userid = ?", id).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, "username = ?", username).Error
	return &user, err
}

func (u *userSQLRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := u.db.Find(&users).Error
	return users, err
}
