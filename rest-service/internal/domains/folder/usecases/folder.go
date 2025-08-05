package usecases

import (
	"context"

	sharedModel "github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	"gorm.io/gorm"
)

type Folder struct {
	folderRepo repos.IFolderRepo
	noteRepo   repos.INoteRepo
	txn        transaction.TxnManager
}

func NewFolder(folderRepo repos.IFolderRepo, noteRepo repos.INoteRepo, txn transaction.TxnManager) *Folder {
	return &Folder{
		folderRepo: folderRepo,
		noteRepo:   noteRepo,
		txn:        txn,
	}
}

func (f *Folder) CreateFolder(ctx context.Context, input *model.CreateFolderInput) error {
	folder := &sharedModel.Folder{
		Name:    input.FolderName,
		OwnerID: input.UserID,
	}

	err := f.folderRepo.CreateFolder(ctx, folder)
	if err != nil {
		return err
	}

	return nil
}

func (f *Folder) GetFolderDetails(ctx context.Context, input *model.GetFolderDetailsInput) (*sharedModel.Folder, error) {
	folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
	if err != nil {
		return nil, err
	}

	if folder.ID == "" {
		return nil, ErrFolderNotFound
	}

	folderShare, err := f.folderRepo.GetFolderShare(ctx, input.FolderID, input.UserID)
	if err != nil {
		return nil, err
	}

	if folder.OwnerID != input.UserID && folderShare.SharedWithUserID == "" {
		return nil, ErrFolderNotSharedToUser
	}

	return folder, nil
}

func (f *Folder) UpdateFolder(ctx context.Context, input *model.UpdateFolderInput) error {
	folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
	if err != nil {
		return err
	}

	if folder.ID == "" {
		return ErrFolderNotFound
	}

	folderShare, err := f.folderRepo.GetFolderShare(ctx, input.FolderID, input.UserID)
	if err != nil {
		return err
	}

	if folder.OwnerID != input.UserID {
		if folderShare.SharedWithUserID == "" {
			return ErrFolderNotSharedToUser
		}

		if folderShare.AccessType != sharedModel.AccessWrite {
			return ErrNoPermission
		}
	}

	folder.Name = input.FolderName
	err = f.folderRepo.UpdateFolder(ctx, folder)
	return err
}

func (f *Folder) DeleteFolder(ctx context.Context, input *model.DeleteFolderInput) error {
	return f.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
		folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
		if err != nil {
			return err
		}

		if folder.ID == "" {
			return ErrFolderNotFound
		}

		if folder.OwnerID != input.UserID {
			return ErrUserNotFolderOwner
		}

		notesOfFolder, err := f.noteRepo.GetNotesByFolderID(ctx, input.FolderID)
		if err != nil {
			return err
		}

		for _, note := range notesOfFolder {
			err = f.noteRepo.DeleteNote(ctx, note.ID)
			if err != nil {
				return err
			}
		}

		err = f.folderRepo.DeleteFolder(ctx, input.FolderID)
		return err
	})
}

func (f *Folder) ShareFolder(ctx context.Context, input *model.ShareFolderInput) error {
	folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
	if err != nil {
		return err
	}

	if folder.ID == "" {
		return ErrFolderNotFound
	}

	if folder.OwnerID != input.CurUserID {
		return ErrUserNotFolderOwner
	}

	for _, userID := range input.SharedUserIDs {
		_, err := f.folderRepo.GetFolderShare(ctx, input.FolderID, userID)
		if err != nil {
			return err
		}

		newShare := &sharedModel.FolderShare{
			FolderID:         input.FolderID,
			SharedWithUserID: userID,
			AccessType:       sharedModel.ToAccessType(input.AccessType),
		}

		err = f.folderRepo.CreateFolderShare(ctx, newShare)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *Folder) RevokeSharingFolder(ctx context.Context, input *model.RevokeSharingFolderInput) error {
	folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
	if err != nil {
		return err
	}

	if folder.ID == "" {
		return ErrFolderNotFound
	}

	if folder.OwnerID != input.CurUserID {
		return ErrUserNotFolderOwner
	}

	folderShare, err := f.folderRepo.GetFolderShare(ctx, input.FolderID, input.SharedUserID)
	if err != nil {
		return err
	}

	if folderShare.FolderID == "" {
		return ErrFolderNotSharedToUser
	}

	err = f.folderRepo.DeleteFolderShare(ctx, folderShare)
	if err != nil {
		return err
	}

	return nil
}
