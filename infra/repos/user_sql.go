package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/internal/models"
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

// Todo: should pass ctx when querying the database
// For that, we can trace, close follow ctx
// But no need to pass ctx for eveytime, cause sometime, if ctx close, still need to query the database
func (u *userSQLRepo) CreateUser(ctx context.Context, user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userSQLRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (u *userSQLRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, "email = ?", email).Error
	return &user, err
}

func (u *userSQLRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := u.db.Find(&users).Error
	return users, err
}
