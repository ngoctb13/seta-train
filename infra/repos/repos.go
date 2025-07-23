package repos

import (
	"github.com/ngoctb13/seta-train/config"
	"github.com/ngoctb13/seta-train/internal/domains/user/repos"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
	// no need config here
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
