package password

import (
	"assess/auth"
)

// type Service interface {
// 	ShowAllPassowrd(userID string) ([]entity.Password, error)
// 	CreateNewPassowrd(user entity.PasswordInput) (entity.Password, error)
// 	ShowPassowrdByID(userID string) (entity.Password, error)
// 	UpdatePassowrdByID(userID string, dataInput entity.PasswordInput) (entity.Password, error)
// }

type service struct {
	repository Repository
	auth       auth.Service
}

func NewService(repo Repository, auth auth.Service) *service {
	return &service{repo, auth}
}

// func (s *service) ShowAllPassowrd(userID string) ([]entity.Password, error) {
// 	users, err := s.repository.FindByID(userID)

// 	if err != nil {
// 		return users, err
// 	}

// 	return formatUsers, nil
// }
