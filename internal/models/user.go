package models

type User struct {
	UserID       string `gorm:"column:userid;default:uuid_generate_v4()"`
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email"`
	PasswordHash string `gorm:"column:passwordhash"`
	Role         string `gorm:"column:role"`
}
