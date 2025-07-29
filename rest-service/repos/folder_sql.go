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

func (f *folderSQLRepo) GetFolderByID(ctx context.Context, id string) (*model.Folder, error) {
	var folder model.Folder
	err := f.db.Where("id = ?", id).First(&folder).Error
	return &folder, err
}

func (f *folderSQLRepo) UpdateFolder(ctx context.Context, folder *model.Folder) error {
	err := f.db.Save(folder).Error
	return err
}

func (f *folderSQLRepo) DeleteFolder(ctx context.Context, id string) error {
	return f.db.Where("id = ?", id).Delete(&model.Folder{}).Error
}
