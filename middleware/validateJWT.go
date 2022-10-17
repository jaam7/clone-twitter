package middleware

import (
	"github.com/jalamar/clone-twitter/routers"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			routers.JSONError(w, "", http.StatusHTTPVersionNotSupported)
			return
		}
		next.ServeHTTP(w, r)
	}
}
