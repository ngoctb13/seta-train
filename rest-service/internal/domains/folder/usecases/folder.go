package usecases

import (
	"context"
	"errors"

	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	sharedModel "github.com/ngoctb13/seta-train/shared-modules/model"
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

	if folder.OwnerID != input.UserID {
		return nil, errors.New("user is not the owner of the folder")
	}

	return folder, nil
}

func (f *Folder) UpdateFolder(ctx context.Context, input *model.UpdateFolderInput) error {
	folder, err := f.folderRepo.GetFolderByID(ctx, input.FolderID)
	if err != nil {
		return err
	}

	if folder.OwnerID != input.UserID {
		return errors.New("user is not the owner of the folder")
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

		if folder.OwnerID != input.UserID {
			return errors.New("user is not the owner of the folder")
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
