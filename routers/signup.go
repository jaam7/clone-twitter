package routers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/jalamar/clone-twitter/bd"
	"github.com/jalamar/clone-twitter/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		http.Error(w, "error in request", http.StatusBadRequest)
	}

	if len(t.Email) == 0 {
		http.Error(w, "email is required", http.StatusUnprocessableEntity)
	}

	if isValidEmail(t.Email) {
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "password must be greather than 6 characters", http.StatusUnprocessableEntity)
	}

	_, isFounded, _ := bd.CheckAvailabilityEmail(t.Email)
	if isFounded {
		http.Error(w, "email already exists", http.StatusBadRequest)
		return
	}

	_, isInserted, err := bd.InsertUser(t)
	if err != nil {
		http.Error(w, "email already exists", http.StatusBadRequest)
		return
	}

	if isInserted {
		http.Error(w, "error while insert user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}
