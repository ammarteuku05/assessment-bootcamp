package migration

type User struct {
	ID        int `gorm:"primaryKer"`
	FullName  string
	Email     string `gorm:"unique"`
	Password  string
	Passwords []Password `gorm:"foreignKey:UserID"`
}

type Password struct {
	ID int `gorm:"primaryKey"`
}
