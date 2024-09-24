package util

import (
    "encoding/json"
    "net/http"
)

// JSONResponse writes a JSON response to the HTTP response writer
func JSONResponse(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// ErrorResponse writes an error message in JSON format
func ErrorResponse(w http.ResponseWriter, status int, message string) {
    JSONResponse(w, status, map[string]string{"error": message})
}
