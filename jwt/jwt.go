package jwt

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jalamar/clone-twitter/models"
)

func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("ale_jaam")

	payload := jwt.MapClaims{
		"name":       user.Name,
		"last_name":  user.LastName,
		"birth_date": user.BirthDate,
		"email":      user.Email,
		"password":   user.Password,
		"avatar":     user.Avatar,
		"banner":     user.Banner,
		"bio":        user.Bio,
		"location":   user.Location,
		"web_site":   user.WebSite,
		"_id":        user.ID.Hex(),
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	}

	log.Println(payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
