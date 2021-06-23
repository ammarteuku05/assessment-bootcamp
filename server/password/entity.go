package password

import "time"

type Password struct {
	ID       int `gorm:"primaryKey"`
	Website  string
	Pass     string
	UserID   int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time `gorm:"index"`
}

type PasswordInput struct {
	Website string `json:"website"`
	Pass    string `json:"pass"`
}
