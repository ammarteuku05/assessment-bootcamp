package config

import (
	"fmt"
	"log"
	"os"

	"assess/migration"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err    = godotenv.Load()
	userDB = os.Getenv("DB_USERNAME")
	passDB = os.Getenv("DB_PASSWORD")
	hostDB = os.Getenv("DB_HOST")
	nameDB = os.Getenv("DB_NAME")
)

func Config() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userDB, passDB, hostDB, nameDB)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&migration.User{}, &migration.Password{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
