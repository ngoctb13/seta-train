package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/shared-modules/model"
	"gorm.io/gorm"
)

type teamSQLRepo struct {
	db *gorm.DB
}

func NewTeamSQLRepo(db *gorm.DB) *teamSQLRepo {
	return &teamSQLRepo{db: db}
}

func (t *teamSQLRepo) CreateTeam(ctx context.Context, team *model.Team) (*model.Team, error) {
	err := t.db.Create(team).Error
	return team, err
}
func (t *teamSQLRepo) GetTeamByID(ctx context.Context, id string) (*model.Team, error) {
	var team model.Team
	err := t.db.First(&team, "id = ?", id).Error
	return &team, err
}
func (t *teamSQLRepo) GetAllTeams(ctx context.Context) ([]*model.Team, error) {
	var teams []*model.Team
	err := t.db.Find(&teams).Error
	return teams, err
}

func (t *teamSQLRepo) AddTeamManager(ctx context.Context, teamManager *model.TeamManager) error {
	err := t.db.Create(teamManager).Error
	return err
}

func (t *teamSQLRepo) RemoveTeamManager(ctx context.Context, teamManager *model.TeamManager) error {
	err := t.db.Delete(teamManager).Error
	return err
}

func (t *teamSQLRepo) IsMainUserManager(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&model.TeamManager{}).
		Where("team_id = ? AND user_id = ? AND is_main_manager = ?", teamID, userID, true).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) IsUserManager(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&model.TeamManager{}).
		Where("team_id = ? AND user_id = ?", teamID, userID).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) IsUserMember(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&model.TeamMember{}).
		Where("team_id = ? AND user_id = ?", teamID, userID).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) AddTeamMember(ctx context.Context, teamMember *model.TeamMember) error {
	err := t.db.Create(teamMember).Error
	return err
}

func (t *teamSQLRepo) RemoveTeamMember(ctx context.Context, teamMember *model.TeamMember) error {
	err := t.db.Delete(teamMember).Error
	return err
}
