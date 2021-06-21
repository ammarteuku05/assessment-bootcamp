package handler

import (
	"assess/auth"
	"assess/entity"
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

func (h *userHandler) GetAllUserHandler(c *gin.Context) {
	users, err := h.userService.ShowAllUser()

	if err != nil {
		errResponse := helper.ResponseAPI("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errResponse)
		return
	}

	respon := helper.ResponseAPI("success", 200, "success", users)

	c.JSON(200, respon)
}

func (h *userHandler) SaveNewUserHandler(c *gin.Context) {
	var input entity.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data required", 400, "bad request", gin.H{"errors": errSplit})

		c.JSON(400, errRespon)
		return
	}

	addUser, err := h.userService.CreateNewUser(input)

	if err != nil {
		errRespon := helper.ResponseAPI("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errRespon)
		return
	}

	res := helper.ResponseAPI("success", 201, "status created", addUser)

	c.JSON(201, res)
}
