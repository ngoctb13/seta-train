package model

type CreateFolderInput struct {
	FolderName string `json:"folder_name"`
	UserID     string `json:"user_id"`
}

type GetFolderDetailsInput struct {
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`
}

type UpdateFolderInput struct {
	FolderID   string `json:"folder_id"`
	FolderName string `json:"folder_name"`
	UserID     string `json:"user_id"`
}

type DeleteFolderInput struct {
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`
}
