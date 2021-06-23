package password

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Password, error)
	Create(user Password) (Password, error)
	FindByID(ID string) (Password, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (Password, error)
	FindByUserId(UserID string) ([]Password, error)
	Delete(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Password, error) {
	var users []Password

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user Password) (Password, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID string) (Password, error) {
	var user Password

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByUserId(UserID string) ([]Password, error) {
	var pass []Password

	if err := r.db.Where("user_id=?", UserID).Find(&pass).Error; err != nil {
		return pass, err
	}

	return pass, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (Password, error) {
	var user Password

	if err := r.db.Model(&user).Where("id=?", ID).Updates(&dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id=?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(ID string) (string, error) {
	var pass Password

	if err := r.db.Where("id = ?", ID).Delete(&pass).Error; err != nil {
		return "error", err
	}

	return "succcess", nil
}
