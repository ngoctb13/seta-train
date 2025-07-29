package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type INoteRepo interface {
	GetNotesByFolderID(ctx context.Context, folderID string) ([]*model.Note, error)
	DeleteNote(ctx context.Context, noteID string) error
}
