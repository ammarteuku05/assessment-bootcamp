package entity

type User struct {
	ID        int `gorm:"primaryKer"`
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	Passwords []Password `gorm:"foreignKey:UserID"`
}

type UserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
