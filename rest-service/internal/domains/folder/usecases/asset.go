package usecases

import (
	"context"

	sharedModel "github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	teamRepos "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	teamUsecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
)

type Asset struct {
	folderRepo repos.IFolderRepo
	noteRepo   repos.INoteRepo
	teamRepo   teamRepos.ITeamRepo
	txn        transaction.TxnManager
}

func NewAsset(folderRepo repos.IFolderRepo, noteRepo repos.INoteRepo, teamRepo teamRepos.ITeamRepo, txn transaction.TxnManager) *Asset {
	return &Asset{
		folderRepo: folderRepo,
		noteRepo:   noteRepo,
		teamRepo:   teamRepo,
		txn:        txn,
	}
}

func (a *Asset) GetAssets(ctx context.Context, input *model.GetAssetsInput) (*sharedModel.TeamAsset, error) {
	team, err := a.teamRepo.GetTeamByID(ctx, input.TeamID)
	if err != nil {
		return nil, err
	}

	if team.ID == "" {
		return nil, teamUsecases.ErrTeamNotFound
	}

	isManager, err := a.teamRepo.IsUserManager(ctx, input.TeamID, input.UserID)
	if err != nil {
		return nil, err
	}

	if !isManager {
		return nil, teamUsecases.ErrUserNotManager
	}

	teamAsset, err := a.folderRepo.GetFoldersByTeamID(ctx, input.TeamID)
	if err != nil {
		return nil, err
	}

	return teamAsset, nil
}
