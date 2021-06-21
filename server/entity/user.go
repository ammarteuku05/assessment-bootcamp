package entity

import "time"

type User struct {
	ID        int `gorm:"primaryKer"`
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	Address   string
	Passwords []Password `gorm:"foreignKey:UserID"`
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time `gorm:"index"`
}

type UserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UpdateUser struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
