package handler

import (
	"assess/auth"
	"assess/helper"
	"assess/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" || len(auth) == 0 {
			errRespon := helper.ResponseAPI("Unauthorization", 401, "error", gin.H{"error": "unauthorize user"})

			c.AbortWithStatusJSON(401, errRespon)
			return
		}

		token, err := authService.ValidateToken(auth)
		if err != nil {
			errRespon := helper.ResponseAPI("Unauthorize", 401, "error", gin.H{"error": err.Error()})
			c.AbortWithStatusJSON(401, errRespon)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			errRespon := helper.ResponseAPI("Unauthorization", 401, "error", gin.H{"error": "unauthorize user"})

			c.AbortWithStatusJSON(401, errRespon)
			return
		}

		userID := int(claim["user_id"].(float64))

		c.Set("currentUser", userID)

	}
}
