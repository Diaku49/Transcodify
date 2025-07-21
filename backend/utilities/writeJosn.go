package utilities

import (
	"encoding/json"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, message string, status int, err error) {
	errorMsg := message + ": " + err.Error()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": errorMsg})
}

func WriteJsonSuccess(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(resp)
}
