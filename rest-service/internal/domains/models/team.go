package model

type CreateTeamInput struct {
	TeamName string `json:"team_name"`
	UserID   string `json:"user_id"`
}
