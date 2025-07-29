package models

type CreateFolderRequest struct {
	FolderName string `json:"folder_name"`
	UserID     string `json:"user_id"`
}