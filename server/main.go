package main

import (
	"assess/auth"
	"assess/config"
	"assess/handler"
	"assess/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB = config.Config()
	userRepo            = user.NewRepository(DB)
	authSer             = auth.NewService()
	userSer             = user.NewService(userRepo, authSer)
	middleware          = handler.Middleware(userSer, authSer)
)

func main() {
	r := gin.Default()

	r.Run()
}
