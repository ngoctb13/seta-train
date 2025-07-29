package usecases

import "errors"

var (
	ErrNoteNotFound       = errors.New("note not found")
	ErrFolderNotFound     = errors.New("folder not found")
	ErrUserNotNoteOwner   = errors.New("user is not the owner of the note")
	ErrUserNotFolderOwner = errors.New("user is not the owner of the folder")
	ErrUserNotMainManager = errors.New("user is not the main manager of the team")
)
