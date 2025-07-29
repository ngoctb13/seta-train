package repos

import (
	folderRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	teamRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
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

func (r *Repo) Teams() teamRepo.ITeamRepo {
	return NewTeamSQLRepo(r.db)
}

func (r *Repo) Folders() folderRepo.IFolderRepo {
	return NewFolderSQLRepo(r.db)
}
