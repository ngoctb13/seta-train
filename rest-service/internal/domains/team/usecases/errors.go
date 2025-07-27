package usecases

import "errors"

var (
	ErrUserNotMainManager = errors.New("user is not main manager of this team")
	ErrUserAlreadyMember  = errors.New("user is already a member of this team")
	ErrUserAlreadyManager = errors.New("user is already a manager of this team")
)
