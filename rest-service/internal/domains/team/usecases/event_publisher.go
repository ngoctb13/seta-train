package usecases

type TeamEventPublisher interface {
	TeamCreated(teamID, teamName, userID string) error
}
