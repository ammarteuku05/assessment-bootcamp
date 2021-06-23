package handler

import (
	"assess/auth"
	"assess/helper"
	"assess/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser user.RegisterInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APINewResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.CreateNewUser(inputUser)
	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APINewResponse(201, "Create new user succeed", newUser)
	c.JSON(201, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser user.InputLogin

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APINewResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APINewResponse(401, "Input data error", gin.H{"errors": err.Error()})

		c.JSON(401, responseError)
		return
	}

	// token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}
	response := helper.APINewResponse(200, "Login user succeed", userData)
	c.JSON(200, response)
}
