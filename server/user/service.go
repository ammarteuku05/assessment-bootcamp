package user

import (
	"assess/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ShowAllUser() ([]UserFormat, error)
	CreateNewUser(user entity.UserInput) (UserFormat, error)
	ShowUserByID(userID string) (UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUser)
	LoginUser(input entity.LoginUser) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) ShowAllUser() ([]UserFormat, error) {
	users, err := s.repository.FindAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) CreateNewUser(user entity.UserInput) (UserFormat, error) {
	genPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var addNew = entity.User{
		FullName: user.FullName,
		Address:  user.Address,
		Email:    user.Email,
		Password: string(genPass),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	addUser, err := s.repository.Create(addNew)
	formatUser := FormatUser(addUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}
