package usecases

import (
	"context"

	model "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	sharedModel "github.com/ngoctb13/seta-train/shared-modules/model"
)

type Team struct {
	teamRepo repos.ITeamRepo
}

func NewTeam(teamRepo repos.ITeamRepo) *Team {
	return &Team{
		teamRepo: teamRepo,
	}
}

func (t *Team) CreateTeam(ctx context.Context, input *model.CreateTeamInput) error {
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
}
