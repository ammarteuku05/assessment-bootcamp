package handler

import (
	"assess/entity"
	"assess/helper"
	"assess/password"
	"assess/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type passHandler struct {
	service password.Service
}

func NewPassHandler(service password.Service) *passHandler {
	return &passHandler{service}
}

func (h *passHandler) GetAllPassbyUser(c *gin.Context) {
	userID := c.MustGet("currenUser").(user.UserFormat)

	userIDtoStr := strconv.Itoa(userID.ID)

	pass, err := h.service.ShowAllPassoword(userIDtoStr)

	if err != nil {
		errResponse := helper.ResponseAPI("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errResponse)
		return
	}

	if userID.ID == 0 {
		errRespon := helper.ResponseAPI("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, errRespon)
		return
	}
	res := helper.ResponseAPI("success", 201, "status created", pass)

	c.JSON(201, res)
}

func (h *passHandler) CreatePassHandler(c *gin.Context) {
	var input entity.PasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data required", 400, "bad request", gin.H{"errors": errSplit})

		c.JSON(400, errRespon)
		return
	}

	userID := c.MustGet("currentUser").(user.UserFormat)

	strconv.Itoa(userID.ID)

	pass, err := h.service.CreateNewPassoword(userID.ID, input)

	if err != nil {
		errRespon := helper.ResponseAPI("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errRespon)
		return
	}

	res := helper.ResponseAPI("success", 201, "status created", pass)

	c.JSON(201, res)
}
func (h *passHandler) ShowByIDPass(c *gin.Context) {
	var input entity.Password

	if err := c.ShouldBindJSON(&input); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data required", 400, "bad request", gin.H{"errors": errSplit})

		c.JSON(400, errRespon)
		return
	}
}
