package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type INoteRepo interface {
	CreateNote(ctx context.Context, note *model.Note) error
	GetNoteByID(ctx context.Context, noteID string) (*model.Note, error)
	GetNotesByFolderID(ctx context.Context, folderID string) ([]*model.Note, error)
	DeleteNote(ctx context.Context, noteID string) error
	UpdateNote(ctx context.Context, note *model.Note) error
}
