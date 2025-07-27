package model

type Team struct {
	ID   string `gorm:"column:id;default:uuid_generate_v4()"`
	Name string `gorm:"column:name"`
}

type TeamMember struct {
	TeamID string `gorm:"column:team_id"`
	UserID string `gorm:"column:user_id"`
}

type TeamManager struct {
	TeamID        string `gorm:"column:team_id"`
	UserID        string `gorm:"column:user_id"`
	IsMainManager bool   `gorm:"column:is_main_manager"`
}
