package models

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	Role         string
}
