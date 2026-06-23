package auth

import (
	"errors"
	"fmt"

	"github.com/yourusername/track-the-trails/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(req RegisterRequest) (string, error) {

	existingUser, err := GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("failed to check existing user")
	}

	if existingUser != nil {
		return "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user := User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Theme:        "light",
	}

	err = CreateUser(&user)
	if err != nil {
		return "", fmt.Errorf("create user error: %v", err)
	}

	token, err := utils.GenerateToken(
		user.ID.String(),
		user.Email,
	)

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func Login(req LoginRequest) (*User, string, error) {

	user, err := GetUserByEmail(req.Email)

	if err != nil {
		return nil, "", errors.New("failed to fetch user")
	}

	if user == nil {
		return nil, "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(
		user.ID.String(),
		user.Email,
	)

	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func GetProfile(userID string) (*User, error) {

	user, err := GetUserByID(userID)

	if err != nil {
		return nil, errors.New("failed to fetch user")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
