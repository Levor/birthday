package models

type User struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"isLoggedIn"`
}
