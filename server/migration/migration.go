package migration

import "time"

type User struct {
	ID        int `gorm:"primaryKer"`
	FullName  string
	Email     string `gorm:"unique"`
	Address   string
	Password  string
	Passwords []Password `gorm:"foreignKey:UserID"`
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time `gorm:"index"`
}

type Password struct {
	ID       int `gorm:"primaryKey"`
	Website  string
	Pass     string
	UserID   int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time `gorm:"index"`
}
