package models

type Worker struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	FIO   string `json:"fio"`
	Day   int    `json:"day"`
	Month int    `json:"month"`
}
