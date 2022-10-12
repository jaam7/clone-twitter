package middleware

import (
	"net/http"

	"github.com/jalamar/clone-twitter/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnection() {
			http.Error(w, "lost DB conecction", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
