package helper

import (
	"errors"
	"math"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type NewResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APINewResponse(code int, msg string, data interface{}) NewResponse {
	var response = NewResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}

	return response
}

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ValidateIDNumber(ID string) error {
	if num, err := strconv.Atoi(ID); err != nil || num == 0 || math.Signbit(float64(num)) == true {
		return errors.New("input must be a valid id user")
	}

	return nil
}
