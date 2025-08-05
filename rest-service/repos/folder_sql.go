package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"gorm.io/gorm"
)

type folderSQLRepo struct {
	db *gorm.DB
}

func NewFolderSQLRepo(db *gorm.DB) *folderSQLRepo {
	return &folderSQLRepo{db: db}
}

func (f *folderSQLRepo) CreateFolder(ctx context.Context, folder *models.Folder) error {
	err := f.db.Create(folder).Error
	return err
}

func (f *folderSQLRepo) GetFolderByID(ctx context.Context, id string) (*models.Folder, error) {
	var folder models.Folder
	err := f.db.Where("id = ?", id).First(&folder).Error
	return &folder, err
}

func (f *folderSQLRepo) UpdateFolder(ctx context.Context, folder *models.Folder) error {
	err := f.db.Save(folder).Error
	return err
}

func (f *folderSQLRepo) DeleteFolder(ctx context.Context, id string) error {
	return f.db.Where("id = ?", id).Delete(&models.Folder{}).Error
}

func (f *folderSQLRepo) GetFolderShare(ctx context.Context, folderID string, userID string) (*models.FolderShare, error) {
	var folderShare models.FolderShare
	err := f.db.Where("folder_id = ? AND shared_with_user_id = ?", folderID, userID).Find(&folderShare).Error
	return &folderShare, err
}

func (f *folderSQLRepo) CreateFolderShare(ctx context.Context, share *models.FolderShare) error {
	err := f.db.Table("folder_shares").Create(share).Error
	return err
}

func (f *folderSQLRepo) DeleteFolderShare(ctx context.Context, share *models.FolderShare) error {
	err := f.db.Where("folder_id = ? AND shared_with_user_id = ?", share.FolderID, share.SharedWithUserID).Delete(&models.FolderShare{}).Error
	return err
}

func (f *folderSQLRepo) GetFoldersByTeamID(ctx context.Context, teamID string) (*models.TeamAsset, error) {
	// Get all folders that team members own or can access
	var folders []*models.TeamAssetFolder
	var team models.Team

	err := f.db.Where("id = ?", teamID).First(&team).Error
	if err != nil {
		return nil, err
	}

	// Query folders owned by team members
	var ownedFolders []struct {
		ID        string `gorm:"column:id"`
		Name      string `gorm:"column:name"`
		OwnerID   string `gorm:"column:owner_id"`
		OwnerName string `gorm:"column:owner_name"`
	}

	err = f.db.Table("folders").
		Select("folders.id, folders.name, folders.owner_id, users.username as owner_name").
		Joins("JOIN users ON folders.owner_id = users.userid").
		Joins("JOIN team_members ON users.userid = team_members.user_id").
		Where("team_members.team_id = ?", teamID).
		Find(&ownedFolders).Error

	if err != nil {
		return nil, err
	}

	// Convert owned folders to TeamAssetFolder
	for _, folder := range ownedFolders {
		teamAssetFolder := &models.TeamAssetFolder{
			ID:        folder.ID,
			Name:      folder.Name,
			OwnerID:   folder.OwnerID,
			OwnerName: folder.OwnerName,
			Notes:     []models.TeamAssetNote{},
		}
		folders = append(folders, teamAssetFolder)
	}

	// Get notes for each folder
	for _, folder := range folders {
		var notes []struct {
			ID    string `gorm:"column:id"`
			Title string `gorm:"column:title"`
			Body  string `gorm:"column:body"`
		}

		// Get notes in this folder
		err = f.db.Table("notes").
			Select("notes.id, notes.title, notes.body").
			Where("notes.folder_id = ?", folder.ID).
			Find(&notes).Error

		if err != nil {
			return nil, err
		}

		// Convert notes to TeamAssetNote
		for _, note := range notes {
			teamAssetNote := models.TeamAssetNote{
				ID:    note.ID,
				Title: note.Title,
				Body:  note.Body,
			}
			folder.Notes = append(folder.Notes, teamAssetNote)
		}
	}

	// Create TeamAsset response
	teamAsset := &models.TeamAsset{
		TeamID:   team.ID,
		TeamName: team.Name,
		Folders:  []models.TeamAssetFolder{},
	}

	for _, folder := range folders {
		teamAsset.Folders = append(teamAsset.Folders, *folder)
	}

	return teamAsset, nil
}
