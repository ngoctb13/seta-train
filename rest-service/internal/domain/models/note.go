package models

type Note struct {
	ID       string `gorm:"column:id;default:uuid_generate_v4()"`
	FolderID string `gorm:"column:folder_id"`
	Title    string `gorm:"column:title"`
	Body     string `gorm:"column:body"`
}

type NoteShare struct {
	NoteID           string     `gorm:"column:note_id"`
	SharedWithUserID string     `gorm:"column:shared_with_user_id"`
	AccessType       AccessType `gorm:"column:access_type"`
}
