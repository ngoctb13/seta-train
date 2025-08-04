package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
)

type ITeamRepo interface {
	CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeamByID(ctx context.Context, id string) (*models.Team, error)
	GetAllTeams(ctx context.Context) ([]*models.Team, error)
	AddTeamManager(ctx context.Context, teamManager *models.TeamManager) error
	RemoveTeamManager(ctx context.Context, teamManager *models.TeamManager) error
	IsMainUserManager(ctx context.Context, teamID, userID string) (bool, error)
	IsUserManager(ctx context.Context, teamID, userID string) (bool, error)
	IsUserMember(ctx context.Context, teamID, userID string) (bool, error)
	AddTeamMember(ctx context.Context, teamMember *models.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamMember *models.TeamMember) error
}
