package models

type CreateTeamReqeust struct {
	TeamName string `json:"team_name"`
}

type AddTeamMembersRequest struct {
	UserIDs []string `json:"user_ids"`
}

type AddTeamManagersRequest struct {
	UserIDs []string `json:"user_ids"`
}
