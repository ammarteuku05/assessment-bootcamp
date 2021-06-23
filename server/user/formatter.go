package user

import (
	"time"
)

type UserFormat struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

type UserLoginFormatter struct {
	ID            int    `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Authorization string `json:"authorization"`
}

func FormatUser(user User) UserFormat {
	var format = UserFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Address:  user.Address,
	}

	return format
}

func UserLoginFormat(user User, token string) UserLoginFormatter {
	return UserLoginFormatter{
		ID:            user.ID,
		FullName:      user.FullName,
		Email:         user.Email,
		Address:       user.Address,
		Authorization: token,
	}
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatDelete(pesan string) DeleteFormat {
	var delete = DeleteFormat{
		Message:    pesan,
		TimeDelete: time.Now(),
	}
	return delete
}
