package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jalamar/clone-twitter/bd"
	"github.com/jalamar/clone-twitter/jwt"
	"github.com/jalamar/clone-twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		JSONError(w, "invalid user/password"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		JSONError(w, "email is required", http.StatusBadRequest)
		return
	}

	userLogin, err := bd.Login(user.Email, user.Password)
	if err != nil {
		JSONError(w, err, http.StatusNotFound)
		return
	}

	jwtKey, err := jwt.GenerateJWT(userLogin)
	if err != nil {
		JSONError(w, "error generating token", http.StatusInternalServerError)
	}

	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
