package model

type GetAssetsInput struct {
	TeamID string
	UserID string
}

type GetAssetsOutput struct {
	TeamID   string
	TeamName string
	Assets   Asset
}

type Asset struct {
	Folders []AssetFolder
}

type AssetFolder struct {
	ID         string
	Name       string
	OwnerID    string
	OwnerName  string
	AccessType string
	Notes      []AssetNote
}

type AssetNote struct {
	ID    string
	Title string
	Body  string
}
