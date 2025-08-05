package usecases

import "errors"

var (
	ErrNoteNotFound          = errors.New("note not found")
	ErrFolderNotFound        = errors.New("folder not found")
	ErrUserNotNoteOwner      = errors.New("user is not the owner of the note")
	ErrUserNotFolderOwner    = errors.New("user is not the owner of the folder")
	ErrFolderNotSharedToUser = errors.New("folder is not shared to this user")
	ErrNoteNotSharedToUser   = errors.New("note is not shared to this user")
	ErrCannotAccessNote      = errors.New("cannot access note")
	ErrNoPermission          = errors.New("do not have permission")
)
