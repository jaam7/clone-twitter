package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	messageResponse := Response{
		StatusCode: code,
		Message:    fmt.Sprintf("%v", err),
	}
	json.NewEncoder(w).Encode(messageResponse)
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
