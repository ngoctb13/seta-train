package usecases

import "errors"

var (
	ErrUserNotMainManager = errors.New("user is not main manager of this team")
	ErrUserNotManager     = errors.New("user is not a manager of this team")
	ErrUserAlreadyMember  = errors.New("user is already a member of this team")
	ErrUserAlreadyManager = errors.New("user is already a manager of this team")
	ErrUserNotMember      = errors.New("user is not a member of this team")
	ErrTeamNotFound       = errors.New("team not found")
)
