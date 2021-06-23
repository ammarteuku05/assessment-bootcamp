package main

import (
	"assess/auth"
	"assess/config"
	"assess/handler"
	"assess/password"
	"assess/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Config()

	// user endpoit
	userRepo    = user.NewRepository(DB)
	authSer     = auth.NewService()
	userSer     = user.NewService(userRepo, authSer)
	middleware  = handler.Middleware(userSer, authSer)
	userHandler = handler.NewUserHandler(userSer, authSer)

	// Password endpoint

	passRepo   = password.NewRepository(DB)
	passSer    = password.NewService(passRepo)
	passHandle = handler.NewPassHandler(passSer)
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

	// user
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)

	// pass

	r.GET("/pass", middleware, passHandle.GetAllPassbyUser)
	r.GET("/pass/:pass_id", middleware, passHandle.ShowByIDPass)
	r.POST("/pass", middleware, passHandle.CreatePassHandler)
	r.PUT("/pass/:pass_id", middleware, passHandle.UpdatebyIDPass)
	r.DELETE("/pass/:pass_id", middleware, passHandle.DeletePassHandler)
	r.Run(":4444")
}
