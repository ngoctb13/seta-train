package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type IFolderRepo interface {
	CreateFolder(ctx context.Context, folder *model.Folder) error
}
