package bd

import "golang.org/x/crypto/bcrypt"

const COST = 8

func EncrytpPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
