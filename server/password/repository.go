package password

import (
	"assess/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Password, error)
	Create(user entity.Password) (entity.Password, error)
	FindByID(ID string) (entity.Password, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Password, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{}
}

func (r *repository) FindAll() ([]entity.Password, error) {
	var users []entity.Password

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entity.Password) (entity.Password, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID string) (entity.Password, error) {
	var user entity.Password

	if err := r.db.Where("user_id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Password, error) {
	var user entity.Password

	if err := r.db.Model(&user).Where("id=?", ID).Updates(&dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id=?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
