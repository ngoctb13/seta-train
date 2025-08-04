package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
)

type IFolderRepo interface {
	CreateFolder(ctx context.Context, folder *models.Folder) error
	GetFolderByID(ctx context.Context, id string) (*models.Folder, error)
	UpdateFolder(ctx context.Context, folder *models.Folder) error
	DeleteFolder(ctx context.Context, id string) error
	GetFolderShare(ctx context.Context, folderID string, userID string) (*models.FolderShare, error)
	CreateFolderShare(ctx context.Context, share *models.FolderShare) error
	DeleteFolderShare(ctx context.Context, share *models.FolderShare) error
	GetFoldersByTeamID(ctx context.Context, teamID string) (*models.TeamAsset, error)
}
