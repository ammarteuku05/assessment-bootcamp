package user

import (
	"assess/auth"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateNewUser(user RegisterInput) (UserFormat, error)
	LoginUser(input InputLogin) (User, error)
}

type service struct {
	repository Repository
	auth       auth.Service
}

func NewService(repo Repository, auth auth.Service) *service {
	return &service{repo, auth}
}

func (s *service) CreateNewUser(user RegisterInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: string(genPassword),
		Address:  user.Address,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) LoginUser(input InputLogin) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newErr := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newErr)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}
