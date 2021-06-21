package user

import (
	"assess/entity"
	"assess/helper"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ShowAllUser() ([]UserFormat, error)
	CreateNewUser(user entity.UserInput) (UserFormat, error)
	ShowUserByID(userID string) (UserFormat, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUser) (UserFormat, error)
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

func (s *service) ShowUserByID(userID string) (UserFormat, error) {
	if err := helper.ValidateID(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newErr := fmt.Sprintf("user id %s not found", userID)

		return UserFormat{}, errors.New(newErr)
	}

	formatUser := FormatUser(user)

	return formatUser, nil
}

func (s *service) UpdateUserByID(userID string, dataInput entity.UpdateUser) (UserFormat, error) {
	var dtUpdate = map[string]interface{}{}

	if err := helper.ValidateID(userID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newErr := fmt.Sprintf("user id %s not found", userID)

		return UserFormat{}, errors.New(newErr)
	}

	dtUpdate["update_at"] = time.Now()

	userUp, err := s.repository.UpdateByID(userID, dtUpdate)

	if err != nil {
		return UserFormat{}, err
	}

	format := FormatUser(userUp)

	return format, nil
}

func (s *service) LoginUser(input entity.LoginUser) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newErr := fmt.Sprintf("user id %s not found", user.ID)
		return user, errors.New(newErr)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}
