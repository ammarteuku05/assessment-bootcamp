package entity

type Password struct {
	ID      int `gorm:"primaryKey"`
	Website string
	Pass    string
	UserID  int
}

type PasswordInput struct {
	Website string `json:"website" binding:"required"`
	Pass    string `json:"pass" binding:"required"`
}
