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

func (f *folderSQLRepo) GetFolderShare(ctx context.Context, folderID string, userID string) (*model.FolderShare, error) {
	var folderShare model.FolderShare
	err := f.db.Where("folder_id = ? AND shared_with_user_id = ?", folderID, userID).Find(&folderShare).Error
	return &folderShare, err
}

func (f *folderSQLRepo) CreateFolderShare(ctx context.Context, share *model.FolderShare) error {
	err := f.db.Create(share).Error
	return err
}

func (f *folderSQLRepo) DeleteFolderShare(ctx context.Context, share *model.FolderShare) error {
	err := f.db.Where("folder_id = ? AND shared_with_user_id = ?", share.FolderID, share.SharedWithUserID).Delete(&model.FolderShare{}).Error
	return err
}
