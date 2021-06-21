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
	DB          *gorm.DB = config.Config()
	userRepo             = user.NewRepository(DB)
	authSer              = auth.NewService()
	userSer              = user.NewService(userRepo, authSer)
	middleware           = handler.Middleware(userSer, authSer)
	userHandler          = handler.NewUserHandler(userSer, authSer)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/users", middleware, userHandler.GetAllUserHandler)
	r.GET("/users/:user_id", middleware, userHandler.GetUserByIDHandler)
	r.POST("/users/register", userHandler.SaveNewUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.PUT("/users/:user_id", middleware, userHandler.UpdateUserHandler)
	r.Run()
}
