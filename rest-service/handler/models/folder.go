package models

type CreateFolderRequest struct {
	FolderName string `json:"folder_name"`
	UserID     string `json:"user_id"`
}

type UpdateFolderRequest struct {
	FolderName string `json:"folder_name"`
}

type ShareFolderRequest struct {
	UserIDs    []string `json:"user_ids"`
	AccessType string   `json:"access_type"`
}
