package usecases

import "errors"

var (
	ErrUserCannotBeNil       = errors.New("user cannot be nil")
	ErrUsernameRequired      = errors.New("username is required")
	ErrUsernameTooShort      = errors.New("username must be at least 3 characters long")
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("invalid email format")
	ErrPasswordRequired      = errors.New("password is required")
	ErrInvalidRole           = errors.New("invalid role, must be MANAGER or MEMBER")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUserIDRequired        = errors.New("user ID is required")
)
