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

func FormatUser(user User) UserFormat {
	var format = UserFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Address:  user.Address,
	}

	return format
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
