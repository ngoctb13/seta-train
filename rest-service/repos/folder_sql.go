package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
	"gorm.io/gorm"
)

type folderSQLRepo struct {
	db *gorm.DB
}

func NewFolderSQLRepo(db *gorm.DB) *folderSQLRepo {
	return &folderSQLRepo{db: db}
}

func (f *folderSQLRepo) CreateFolder(ctx context.Context, folder *model.Folder) error {
	err := f.db.Create(folder).Error
	return err
}
