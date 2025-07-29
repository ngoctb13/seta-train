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

type ShareFolderInput struct {
	FolderID      string   `json:"folder_id"`
	CurUserID     string   `json:"cur_user_id"`
	SharedUserIDs []string `json:"shared_user_ids"`
	AccessType    string   `json:"access_type"`
}

type RevokeSharingFolderInput struct {
	CurUserID    string `json:"cur_user_id"`
	FolderID     string `json:"folder_id"`
	SharedUserID string `json:"shared_user_id"`
}
