package user

import (
	"assess/auth"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateNewUser(user RegisterInput) (UserFormat, error)
	LoginUser(input InputLogin) (UserLoginFormatter, error)
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

func (s *service) LoginUser(input InputLogin) (UserLoginFormatter, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return UserLoginFormatter{}, err
	}

	if user.ID == 0 || len(user.FullName) <= 1 {
		return UserLoginFormatter{}, errors.New("user email / password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return UserLoginFormatter{}, errors.New("user email / password invalid")
	}

	token, _ := s.auth.GenerateToken(user.ID)

	formatter := UserLoginFormat(user, token)

	return formatter, nil
}
