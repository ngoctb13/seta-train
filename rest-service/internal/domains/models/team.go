package model

type CreateTeamInput struct {
	TeamName string `json:"team_name"`
	UserID   string `json:"user_id"`
}

type AddTeamMembersInput struct {
	CurUserID string   `json:"cur_user_id"`
	TeamID    string   `json:"team_id"`
	UserIDs   []string `json:"user_ids"`
}

type AddTeamManagersInput struct {
	CurUserID string   `json:"cur_user_id"`
	TeamID    string   `json:"team_id"`
	UserIDs   []string `json:"user_ids"`
}
