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

type RemoveTeamMemberInput struct {
	CurUserID string `json:"cur_user_id"`
	TeamID    string `json:"team_id"`
	MemberID  string `json:"member_id"`
}

type RemoveTeamManagerInput struct {
	CurUserID string `json:"cur_user_id"`
	TeamID    string `json:"team_id"`
	ManagerID string `json:"manager_id"`
}
