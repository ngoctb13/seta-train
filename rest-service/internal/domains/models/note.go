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

type ShareNoteInput struct {
	NoteID        string   `json:"note_id"`
	CurUserID     string   `json:"cur_user_id"`
	SharedUserIDs []string `json:"shared_user_ids"`
	AccessType    string   `json:"access_type"`
}

type RevokeSharingNoteInput struct {
	CurUserID    string `json:"cur_user_id"`
	NoteID       string `json:"note_id"`
	SharedUserID string `json:"shared_user_id"`
}
