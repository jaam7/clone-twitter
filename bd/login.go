package bd

import (
	"errors"

	"github.com/jalamar/clone-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.User, error) {
	user, finded, _ := CheckAvailabilityEmail(email)
	if !finded {
		return models.User{}, errors.New("unregistered user")
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return models.User{}, errors.New("incorrect user/passowrd")
	}

	return user, nil
}
