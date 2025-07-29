package model

type Note struct {
	Title string
	Body  string
}

type CreateNotesInput struct {
	FolderID string
	Notes    []Note
	UserID   string
}

type ViewNoteInput struct {
	UserID string
	NoteID string
}

type UpdateNoteInput struct {
	UserID string
	NoteID string
	Note   Note
}

type DeleteNoteInput struct {
	UserID string
	NoteID string
}
