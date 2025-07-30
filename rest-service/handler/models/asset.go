package models

type AssetNote struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type AssetFolder struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	OwnerID    string      `json:"owner_id"`
	OwnerName  string      `json:"owner_name"`
	AccessType string      `json:"access_type"`
	Notes      []AssetNote `json:"notes"`
}

type Asset struct {
	Folders []AssetFolder `json:"folders"`
}

type GetAssetResponse struct {
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	Assets   Asset  `json:"assets"`
}
