package usecases

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	sharedModel "github.com/ngoctb13/seta-train/shared-modules/model"
)

type Folder struct {
	folderRepo repos.IFolderRepo
	txn        transaction.TxnManager
}

func NewFolder(folderRepo repos.IFolderRepo, txn transaction.TxnManager) *Folder {
	return &Folder{
		folderRepo: folderRepo,
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
