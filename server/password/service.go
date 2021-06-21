package password

import (
	"assess/entity"
	"fmt"
	"time"
)

type Service interface {
	ShowAllPassoword(userID string) ([]entity.Password, error)
	CreateNewPassoword(userID int, user entity.PasswordInput) (entity.Password, error)
	ShowPassowordByID(userID string) (entity.Password, error)
	UpdatePassowordByID(userID string, dataInput entity.PasswordInput) (entity.Password, error)
	DeletePassword(passID string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) ShowAllPassoword(userID string) ([]entity.Password, error) {
	pass, err := s.repository.FindByUserId(userID)

	if err != nil {
		return pass, err
	}

	return pass, nil
}

func (s *service) CreateNewPassoword(userID int, user entity.PasswordInput) (entity.Password, error) {
	var newPass = entity.Password{
		UserID:   userID,
		Website:  user.Website,
		Pass:     user.Pass,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	pass, err := s.repository.Create(newPass)

	if err != nil {
		return pass, err
	}

	return pass, nil
}

func (s *service) ShowPassowordByID(userID string) (entity.Password, error) {
	pass, err := s.repository.FindByID(userID)

	if err != nil {
		return pass, err
	}

	return pass, nil
}

func (s *service) UpdatePassowordByID(userID string, dataInput entity.PasswordInput) (entity.Password, error) {
	var dtUpdate = map[string]interface{}{}

	if dataInput.Pass != "" {
		dtUpdate["password"] = dataInput.Pass
	}

	if dataInput.Website != "" {
		dtUpdate["website"] = dataInput.Website
	}

	dtUpdate["update_at"] = time.Now()

	pass, err := s.repository.UpdateByID(userID, dtUpdate)

	if err != nil {
		return pass, err
	}

	return pass, nil

}

func (s *service) DeletePassword(passID string) (string, error) {
	mess, err := s.repository.Delete(passID)

	if err != nil || mess == "error" {
		return mess, err
	}

	msg := fmt.Sprintf("password id %s success deleted", passID)

	return msg, nil
}
