package bd

import (
	"github.com/jalamar/clone-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.User, bool){
	user, finded, _:= CheckAvailabilityEmail(email)
	if finded {
		return models.User{}, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return models.User{}, false
	}

	return user, true
}