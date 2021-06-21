package helper

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"Message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseAPI(mess string, code int, status string, data interface{}) Response {
	var meta = Meta{
		Message: mess,
		Code:    code,
		Status:  status,
	}

	var response = Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func DivErrInfor(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ValidateID(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 {
		return errors.New("Input have to be valid user id")

	}
	return nil
}
