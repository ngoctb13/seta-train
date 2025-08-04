package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
)

type INoteRepo interface {
	CreateNote(ctx context.Context, note *models.Note) error
	GetNoteByID(ctx context.Context, noteID string) (*models.Note, error)
	GetNotesByFolderID(ctx context.Context, folderID string) ([]*models.Note, error)
	DeleteNote(ctx context.Context, noteID string) error
	UpdateNote(ctx context.Context, note *models.Note) error
	GetNoteShare(ctx context.Context, noteID string, userID string) (*models.NoteShare, error)
	CreateNoteShare(ctx context.Context, share *models.NoteShare) error
	DeleteNoteShare(ctx context.Context, share *models.NoteShare) error
}
