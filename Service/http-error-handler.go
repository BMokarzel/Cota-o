package service

import (
	"encoding/json"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, code int, v any) error {

	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)

}
