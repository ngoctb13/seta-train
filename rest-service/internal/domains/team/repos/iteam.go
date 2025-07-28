package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
)

type ITeamRepo interface {
	CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error)
	GetTeamByID(ctx context.Context, id string) (*model.Team, error)
	GetAllTeams(ctx context.Context) ([]*model.Team, error)
	AddTeamManager(ctx context.Context, teamManager *model.TeamManager) error
	RemoveTeamManager(ctx context.Context, teamManager *model.TeamManager) error
	IsMainUserManager(ctx context.Context, teamID, userID string) (bool, error)
	IsUserManager(ctx context.Context, teamID, userID string) (bool, error)
	IsUserMember(ctx context.Context, teamID, userID string) (bool, error)
	AddTeamMember(ctx context.Context, teamMember *model.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamMember *model.TeamMember) error
}
