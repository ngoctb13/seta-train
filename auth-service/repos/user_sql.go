package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/auth-service/internal/domain/models"
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

func (u *userSQLRepo) CreateUser(ctx context.Context, user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userSQLRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "userid = ?", id).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "username = ?", username).Error
	return &user, err
}

func (u *userSQLRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userSQLRepo) UpdateUserRole(ctx context.Context, userID string, role string) error {
	return u.db.Model(&models.User{}).Where("userid = ?", userID).Update("role", role).Error
}
