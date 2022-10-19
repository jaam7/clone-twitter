package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jalamar/clone-twitter/bd"
	"github.com/jalamar/clone-twitter/models"
)

var Email, IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("ale_jaam")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid format token")
	}

	token = strings.TrimSpace(splitToken[1])

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return claims, false, string(""), err
	}

	_, finded, _ := bd.CheckAvailabilityEmail(claims.Email)
	if !finded {
		return claims, finded, IDUser, errors.New("user not found")
	}

	Email = claims.Email
	IDUser = claims.ID.Hex()

	return claims, finded, IDUser, nil
}
