package entity

import "time"

type Password struct {
	ID       int `gorm:"primaryKey"`
	Website  string
	Pass     string
	UserID   int
	CreateAt time.Time
	UpdateAt time.Time
}

type PasswordInput struct {
	Website string `json:"website" binding:"required"`
	Pass    string `json:"pass" binding:"required"`
}
