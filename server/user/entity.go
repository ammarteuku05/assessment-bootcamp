package user

import (
	"assess/password"
	"time"
)

type User struct {
	ID       int `gorm:"primaryKey"`
	FullName string
	Email    string
	Password string
	Address  string
	Pass     []password.Password `gorm:"foreignKey=UserID"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time `gorm:"index"`
}
