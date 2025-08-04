package usecases

import (
	"context"

	sharedModel "github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
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

	folderShare, err := n.folderRepo.GetFolderShare(ctx, note.FolderID, input.UserID)
	if err != nil {
		return nil, err
	}

	noteShare, err := n.noteRepo.GetNoteShare(ctx, input.NoteID, input.UserID)
	if err != nil {
		return nil, err
	}

	if folder.OwnerID != input.UserID && (folderShare.ID == "" && noteShare.ID == "") {
		return nil, ErrNoteNotSharedToUser
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

	folderShare, err := n.folderRepo.GetFolderShare(ctx, note.FolderID, input.UserID)
	if err != nil {
		return err
	}

	noteShare, err := n.noteRepo.GetNoteShare(ctx, input.NoteID, input.UserID)
	if err != nil {
		return err
	}

	if folder.OwnerID != input.UserID && (folderShare.ID == "" && noteShare.ID == "") {
		return ErrNoteNotSharedToUser
	}

	if noteShare.AccessType != sharedModel.AccessWrite && folderShare.AccessType != sharedModel.AccessWrite {
		return ErrCannotAccessNote
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

func (n *Note) ShareNote(ctx context.Context, input *model.ShareNoteInput) error {
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

	if folder.OwnerID != input.CurUserID {
		return ErrUserNotNoteOwner
	}

	for _, userID := range input.SharedUserIDs {
		noteShare, err := n.noteRepo.GetNoteShare(ctx, input.NoteID, userID)
		if err != nil {
			return err
		}

		if noteShare != nil {
			continue
		}

		newShare := &sharedModel.NoteShare{
			NoteID:           input.NoteID,
			SharedWithUserID: userID,
			AccessType:       sharedModel.ToAccessType(input.AccessType),
		}

		err = n.noteRepo.CreateNoteShare(ctx, newShare)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *Note) RevokeSharingNote(ctx context.Context, input *model.RevokeSharingNoteInput) error {
	return n.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
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

		if folder.OwnerID != input.CurUserID {
			return ErrUserNotNoteOwner
		}

		noteShare, err := n.noteRepo.GetNoteShare(ctx, input.NoteID, input.SharedUserID)
		if err != nil {
			return err
		}

		if noteShare.ID == "" {
			return ErrNoteNotSharedToUser
		}

		err = n.noteRepo.DeleteNoteShare(ctx, noteShare)
		if err != nil {
			return err
		}

		folderShare, err := n.folderRepo.GetFolderShare(ctx, note.FolderID, input.SharedUserID)
		if err != nil {
			return err
		}

		if folderShare.ID == "" {
			return ErrFolderNotSharedToUser
		}

		err = n.folderRepo.DeleteFolderShare(ctx, folderShare)
		if err != nil {
			return err
		}

		return nil
	})
}
