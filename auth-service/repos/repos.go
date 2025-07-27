package repos

import (
	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"gorm.io/gorm"
)

type Repo struct {
	db  *gorm.DB
	cfg *config.PostgresConfig
}

func NewSQLRepo(db *gorm.DB, cfg *config.PostgresConfig) IRepo {
	return &Repo{
		db:  db,
		cfg: cfg,
	}
}

func (r *Repo) Users() repos.IUserRepo {
	return NewUserSQLRepo(r.db)
}
