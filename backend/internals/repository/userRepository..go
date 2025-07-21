package repository

import (
	"fmt"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	err := ur.DB.Create(user).Error
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(payload model.UserLoginPayload) (*model.UserCredentials, error) {
	var userCred model.UserCredentials
	err := ur.DB.Find(&userCred).Where("email = ?", payload.Email).Select("id, password").Error
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	return &userCred, nil
}

func (ur *UserRepository) GetProfileById(id uint) (*model.User, error) {
	var user model.User
	err := ur.DB.Find(&user).Where("user_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("db error: %v", err)
	}

	return &user, nil
}
