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
}
