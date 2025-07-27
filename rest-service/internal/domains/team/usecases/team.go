package usecases

import (
	"context"

	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	"github.com/ngoctb13/seta-train/shared-modules/infra/transaction"
	sharedModel "github.com/ngoctb13/seta-train/shared-modules/model"
	"gorm.io/gorm"
)

type Team struct {
	teamRepo repos.ITeamRepo
	txn      transaction.TxnManager
}

func NewTeam(teamRepo repos.ITeamRepo, txn transaction.TxnManager) *Team {
	return &Team{
		teamRepo: teamRepo,
		txn:      txn,
	}
}

func (t *Team) CreateTeam(ctx context.Context, input *model.CreateTeamInput) error {
	return t.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
		teamInput := sharedModel.Team{
			Name: input.TeamName,
		}

		team, err := t.teamRepo.CreateTeam(ctx, &teamInput)
		if err != nil {
			return err
		}

		teamManagerInput := sharedModel.TeamManager{
			TeamID:        team.ID,
			UserID:        input.UserID,
			IsMainManager: true,
		}

		err = t.teamRepo.AddTeamManager(ctx, &teamManagerInput)
		if err != nil {
			return err
		}

		return nil
	})
}

func (t *Team) AddTeamMembers(ctx context.Context, input *model.AddTeamMembersInput) error {
	return t.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
		isMainManager, err := t.teamRepo.IsMainUserManager(ctx, input.TeamID, input.CurUserID)
		if err != nil {
			return err
		}
		if !isMainManager {
			return ErrUserNotMainManager
		}

		for _, userID := range input.UserIDs {
			isMember, _ := t.teamRepo.IsUserMember(ctx, input.TeamID, userID)
			if isMember {
				return ErrUserAlreadyMember
			}
		}

		for _, userID := range input.UserIDs {
			teamMemberInput := sharedModel.TeamMember{
				TeamID: input.TeamID,
				UserID: userID,
			}

			err = t.teamRepo.AddTeamMember(ctx, &teamMemberInput)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (t *Team) AddTeamManagers(ctx context.Context, input *model.AddTeamManagersInput) error {
	return t.txn.WithTransaction(ctx, func(tx *gorm.DB) error {
		var count int64

		isMainManager, err := t.teamRepo.IsMainUserManager(ctx, input.TeamID, input.CurUserID)
		if err != nil {
			return err
		}

		if !isMainManager {
			return ErrUserNotMainManager
		}

		for _, userID := range input.UserIDs {
			isManager, _ := t.teamRepo.IsUserManager(ctx, input.TeamID, userID)
			if isManager {
				count++
			}
		}

		if count > 0 {
			return ErrUserAlreadyManager
		}

		for _, userID := range input.UserIDs {
			teamManagerInput := sharedModel.TeamManager{
				TeamID:        input.TeamID,
				UserID:        userID,
				IsMainManager: false,
			}

			err = t.teamRepo.AddTeamManager(ctx, &teamManagerInput)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
