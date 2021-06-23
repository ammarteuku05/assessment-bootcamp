package handler

import (
	"assess/helper"
	"assess/password"
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
	userID := int(c.MustGet("currentUser").(int))

	userIDtoStr := strconv.Itoa(userID)

	pass, err := h.service.ShowAllPassoword(userIDtoStr)

	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	if userID == 0 {
		errRespon := helper.APINewResponse(401, "Unauthorize", gin.H{"error": "user ID not authorize"})

		c.JSON(401, errRespon)
		return
	}
	res := helper.APINewResponse(201, "success", pass)

	c.JSON(201, res)
}

func (h *passHandler) CreatePassHandler(c *gin.Context) {
	var input password.PasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APINewResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userID := int(c.MustGet("currentUser").(int))

	pass, err := h.service.CreateNewPassoword(userID, input)

	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	res := helper.APINewResponse(201, "success", pass)

	c.JSON(201, res)
}
func (h *passHandler) ShowByIDPass(c *gin.Context) {
	id := c.Params.ByName("pass_id")

	pass, err := h.service.ShowPassowordByID(id)

	if err != nil {
		errRespon := helper.APINewResponse(500, "error", gin.H{"error": err.Error()})

		c.JSON(500, errRespon)
		return
	}

	res := helper.APINewResponse(201, "status create", pass)
	c.JSON(200, res)
}

func (h *passHandler) UpdatebyIDPass(c *gin.Context) {
	passID := c.Params.ByName("pass_id")

	var input password.PasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APINewResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	pass, err := h.service.UpdatePassowordByID(passID, input)

	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	res := helper.APINewResponse(200, "success", pass)
	c.JSON(200, res)
}

func (h *passHandler) DeletePassHandler(c *gin.Context) {
	passID := c.Params.ByName("pass_id")

	msg, err := h.service.DeletePassword(passID)

	if err != nil {
		responseError := helper.APINewResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	res := helper.APINewResponse(200, "success", msg)
	c.JSON(200, res)
}
