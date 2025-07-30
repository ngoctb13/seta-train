package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type IFolderRepo interface {
	CreateFolder(ctx context.Context, folder *model.Folder) error
	GetFolderByID(ctx context.Context, id string) (*model.Folder, error)
	UpdateFolder(ctx context.Context, folder *model.Folder) error
	DeleteFolder(ctx context.Context, id string) error
	GetFolderShare(ctx context.Context, folderID string, userID string) (*model.FolderShare, error)
	CreateFolderShare(ctx context.Context, share *model.FolderShare) error
	DeleteFolderShare(ctx context.Context, share *model.FolderShare) error
	GetFoldersByTeamID(ctx context.Context, teamID string) (*model.TeamAsset, error)
}
