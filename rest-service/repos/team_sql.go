package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"gorm.io/gorm"
)

type teamSQLRepo struct {
	db *gorm.DB
}

func NewTeamSQLRepo(db *gorm.DB) *teamSQLRepo {
	return &teamSQLRepo{db: db}
}

func (t *teamSQLRepo) CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	err := t.db.Create(team).Error
	return team, err
}
func (t *teamSQLRepo) GetTeamByID(ctx context.Context, id string) (*models.Team, error) {
	var team models.Team
	err := t.db.First(&team, "id = ?", id).Error
	return &team, err
}
func (t *teamSQLRepo) GetAllTeams(ctx context.Context) ([]*models.Team, error) {
	var teams []*models.Team
	err := t.db.Find(&teams).Error
	return teams, err
}

func (t *teamSQLRepo) AddTeamManager(ctx context.Context, teamManager *models.TeamManager) error {
	err := t.db.Create(teamManager).Error
	return err
}

func (t *teamSQLRepo) RemoveTeamManager(ctx context.Context, teamManager *models.TeamManager) error {
	return t.db.Where("team_id = ? AND user_id = ?", teamManager.TeamID, teamManager.UserID).Delete(&models.TeamManager{}).Error
}

func (t *teamSQLRepo) IsMainUserManager(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&models.TeamManager{}).
		Where("team_id = ? AND user_id = ? AND is_main_manager = ?", teamID, userID, true).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) IsUserManager(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&models.TeamManager{}).
		Where("team_id = ? AND user_id = ?", teamID, userID).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) IsUserMember(ctx context.Context, teamID, userID string) (bool, error) {
	var count int64
	err := t.db.Model(&models.TeamMember{}).
		Where("team_id = ? AND user_id = ?", teamID, userID).
		Count(&count).Error
	return count > 0, err
}

func (t *teamSQLRepo) AddTeamMember(ctx context.Context, teamMember *models.TeamMember) error {
	err := t.db.Create(teamMember).Error
	return err
}

func (t *teamSQLRepo) RemoveTeamMember(ctx context.Context, teamMember *models.TeamMember) error {
	return t.db.Where("team_id = ? AND user_id = ?", teamMember.TeamID, teamMember.UserID).Delete(&models.TeamMember{}).Error
}
