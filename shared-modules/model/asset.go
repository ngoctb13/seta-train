package model

type Folder struct {
	ID      string `gorm:"column:id;default:uuid_generate_v4()"`
	Name    string `gorm:"column:name"`
	OwnerID string `gorm:"column:owner_id"`
}

type Note struct {
	ID       string `gorm:"column:id;default:uuid_generate_v4()"`
	FolderID string `gorm:"column:folder_id"`
	Title    string `gorm:"column:title"`
	Body     string `gorm:"column:body"`
}

type AccessType string

const (
	AccessRead  AccessType = "read"
	AccessWrite AccessType = "write"
)

type FolderShare struct {
	ID               string     `gorm:"column:id;default:uuid_generate_v4()"`
	FolderID         string     `gorm:"column:folder_id"`
	SharedWithUserID string     `gorm:"column:shared_with_user_id"`
	AccessType       AccessType `gorm:"column:access_type"`
}

type NoteShare struct {
	ID               string     `gorm:"column:id;default:uuid_generate_v4()"`
	NoteID           string     `gorm:"column:note_id"`
	SharedWithUserID string     `gorm:"column:shared_with_user_id"`
	AccessType       AccessType `gorm:"column:access_type"`
}

func ToAccessType(accessType string) AccessType {
	switch accessType {
	case "read":
		return AccessRead
	case "write":
		return AccessWrite
	default:
		return AccessRead
	}
}

type TeamAsset struct {
	TeamID   string            `json:"teamId"`
	TeamName string            `json:"teamName"`
	Folders  []TeamAssetFolder `json:"folders"`
}

type TeamAssetFolder struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	OwnerID   string          `json:"ownerId"`
	OwnerName string          `json:"ownerName"`
	Notes     []TeamAssetNote `json:"notes"`
}

type TeamAssetNote struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
