package usecases

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	sharedModel "github.com/ngoctb13/seta-train/shared-modules/model"
	"gorm.io/gorm"
)

type Note struct {
	noteRepo   repos.INoteRepo
	folderRepo repos.IFolderRepo
	txn        transaction.TxnManager
}

func NewNote(noteRepo repos.INoteRepo, folderRepo repos.IFolderRepo, txn transaction.TxnManager) *Note {
	return &Note{noteRepo: noteRepo, folderRepo: folderRepo, txn: txn}
}

func (n *Note) CreateNotes(ctx context.Context, input *model.CreateNotesInput) error {
	return n.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
		folder, err := n.folderRepo.GetFolderByID(ctx, input.FolderID)
		if err != nil {
			return err
		}

		if folder.ID == "" {
			return ErrFolderNotFound
		}

		if folder.OwnerID != input.UserID {
			return ErrUserNotFolderOwner
		}

		for _, note := range input.Notes {
			note := &sharedModel.Note{
				FolderID: input.FolderID,
				Title:    note.Title,
				Body:     note.Body,
			}

			err = n.noteRepo.CreateNote(ctx, note)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (n *Note) ViewNote(ctx context.Context, input *model.ViewNoteInput) (*sharedModel.Note, error) {
	note, err := n.noteRepo.GetNoteByID(ctx, input.NoteID)
	if err != nil {
		return nil, err
	}

	if note.ID == "" {
		return nil, ErrNoteNotFound
	}

	folder, err := n.folderRepo.GetFolderByID(ctx, note.FolderID)
	if err != nil {
		return nil, err
	}

	if folder.ID == "" {
		return nil, ErrFolderNotFound
	}

	if folder.OwnerID != input.UserID {
		return nil, ErrUserNotNoteOwner
	}

	return note, nil
}

func (n *Note) UpdateNote(ctx context.Context, input *model.UpdateNoteInput) error {
	note, err := n.noteRepo.GetNoteByID(ctx, input.NoteID)
	if err != nil {
		return err
	}

	if note.ID == "" {
		return ErrNoteNotFound
	}

	folder, err := n.folderRepo.GetFolderByID(ctx, note.FolderID)
	if err != nil {
		return err
	}

	if folder.ID == "" {
		return ErrFolderNotFound
	}

	if folder.OwnerID != input.UserID {
		return ErrUserNotNoteOwner
	}

	note.Title = input.Note.Title
	note.Body = input.Note.Body

	err = n.noteRepo.UpdateNote(ctx, note)
	if err != nil {
		return err
	}

	return nil
}

func (n *Note) DeleteNote(ctx context.Context, input *model.DeleteNoteInput) error {
	note, err := n.noteRepo.GetNoteByID(ctx, input.NoteID)
	if err != nil {
		return err
	}

	if note.ID == "" {
		return ErrNoteNotFound
	}

	folder, err := n.folderRepo.GetFolderByID(ctx, note.FolderID)
	if err != nil {
		return err
	}

	if folder.ID == "" {
		return ErrFolderNotFound
	}

	if folder.OwnerID != input.UserID {
		return ErrUserNotNoteOwner
	}

	err = n.noteRepo.DeleteNote(ctx, input.NoteID)
	if err != nil {
		return err
	}

	return nil
}
