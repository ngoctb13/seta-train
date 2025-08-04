package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"gorm.io/gorm"
)

type noteSQLRepo struct {
	db *gorm.DB
}

func NewNoteSQLRepo(db *gorm.DB) *noteSQLRepo {
	return &noteSQLRepo{db: db}
}

func (n *noteSQLRepo) CreateNote(ctx context.Context, note *models.Note) error {
	return n.db.Create(note).Error
}

func (n *noteSQLRepo) GetNotesByFolderID(ctx context.Context, folderID string) ([]*models.Note, error) {
	var notes []*models.Note
	err := n.db.Where("folder_id = ?", folderID).Find(&notes).Error
	return notes, err
}

func (n *noteSQLRepo) DeleteNote(ctx context.Context, noteID string) error {
	return n.db.Where("id = ?", noteID).Delete(&models.Note{}).Error
}

func (n *noteSQLRepo) GetNoteByID(ctx context.Context, noteID string) (*models.Note, error) {
	var note *models.Note
	err := n.db.Where("id = ?", noteID).First(&note).Error
	return note, err
}

func (n *noteSQLRepo) UpdateNote(ctx context.Context, note *models.Note) error {
	err := n.db.Save(note).Error
	return err
}

func (n *noteSQLRepo) GetNoteShare(ctx context.Context, noteID string, userID string) (*models.NoteShare, error) {
	var noteShare models.NoteShare
	err := n.db.Where("note_id = ? AND shared_with_user_id = ?", noteID, userID).First(&noteShare).Error
	return &noteShare, err
}

func (n *noteSQLRepo) CreateNoteShare(ctx context.Context, share *models.NoteShare) error {
	err := n.db.Create(share).Error
	return err
}

func (n *noteSQLRepo) DeleteNoteShare(ctx context.Context, share *models.NoteShare) error {
	err := n.db.Where("note_id = ? AND shared_with_user_id = ?", share.NoteID, share.SharedWithUserID).Delete(&models.NoteShare{}).Error
	return err
}
