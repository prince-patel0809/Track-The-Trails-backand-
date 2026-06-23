package auth

import (
	"errors"

	"github.com/yourusername/track-the-trails/config"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*User, error) {

	var user User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func CreateUser(user *User) error {

	return config.DB.Create(user).Error
}

func GetUserByID(id string) (*User, error) {

	var user User

	err := config.DB.
		Where("id = ?", id).
		First(&user).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
