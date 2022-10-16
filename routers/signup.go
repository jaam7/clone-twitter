package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/jalamar/clone-twitter/bd"
	"github.com/jalamar/clone-twitter/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		JSONError(w, "error in request", http.StatusBadRequest)
		return
	}

	log.Println(t)

	if len(t.Email) == 0 {
		JSONError(w, "email is required", http.StatusUnprocessableEntity)
		return
	}

	if len(t.Password) == 0 {
		JSONError(w, "password is required", http.StatusUnprocessableEntity)
		return
	}

	if !isValidEmail(t.Email) {
		JSONError(w, "email is invalid", http.StatusUnprocessableEntity)
		return
	}

	log.Println(isValidEmail(t.Email))

	if len(t.Password) < 6 {
		JSONError(w, "password must be greather than 6 characters", http.StatusUnprocessableEntity)
		return
	}

	_, isFounded, _ := bd.CheckAvailabilityEmail(t.Email)
	log.Println(isFounded)

	if isFounded {
		JSONError(w, "email already exists", http.StatusConflict)
		return
	}

	_, isInserted, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "error with DB", http.StatusBadRequest)
		return
	}

	if !isInserted {
		http.Error(w, "error while insert user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}
