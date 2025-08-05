package repository

import (
	"fmt"

	"github.com/Diaku49/FoodOrderSystem/backend/internals/model"
	"golang.org/x/crypto/bcrypt"
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

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var userCred model.User
	err := ur.DB.Where("email = ?", email).Select("id, password, email").Find(&userCred).Error
	if err != nil {
		return nil, fmt.Errorf("db error: %w", err)
	}

	if userCred.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &userCred, nil
}

func (ur *UserRepository) GetProfileById(id uint) (*model.User, error) {
	var user model.User
	err := ur.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("db error: %v", err)
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (ur *UserRepository) ChangePassword(userId uint, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed hashing: %w", err)
	}

	err = ur.DB.Model(&model.User{}).Where("id = ?", userId).Update("password", string(hashedPassword)).Error
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	return nil
}
