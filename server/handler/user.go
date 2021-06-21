package handler

import (
	"assess/auth"
	"assess/entity"
	"assess/helper"
	"assess/user"
	"strconv"

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

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.ShowUserByID(id)

	if err != nil {
		errRespon := helper.ResponseAPI("error bad requeest", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, errRespon)
		return
	}
	res := helper.ResponseAPI("success", 201, "status create", user)
	c.JSON(200, res)
}

func (h *userHandler) UpdateUserHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var update entity.UpdateUser

	if err := c.ShouldBindJSON(&update); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data required", 400, "bad request", gin.H{"errors": errSplit})

		c.JSON(400, errRespon)
		return
	}

	idParam, _ := strconv.Atoi(id)

	datUser := int(c.MustGet("currentUser").(int))

	if idParam != datUser {
		errRespon := helper.ResponseAPI("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, errRespon)
		return
	}

	user, err := h.userService.UpdateUserByID(id, update)
	if err != nil {
		errRespon := helper.ResponseAPI("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errRespon)
		return
	}

	res := helper.ResponseAPI("successs", 200, "success", user)
	c.JSON(200, res)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var input entity.LoginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data required", 400, "bad request", gin.H{"errors": errSplit})

		c.JSON(400, errRespon)
		return
	}

	user, err := h.userService.LoginUser(input)

	if err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input data error", 401, "bad request", gin.H{"errors": errSplit})

		c.JSON(401, errRespon)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)

	if err := c.ShouldBindJSON(&input); err != nil {
		errSplit := helper.DivErrInfor(err)
		errRespon := helper.ResponseAPI("input server error", 500, "error", gin.H{"errors": errSplit})

		c.JSON(401, errRespon)
		return
	}

	res := helper.ResponseAPI("success", 200, "success", gin.H{"toker": token})
	c.JSON(200, res)
}
