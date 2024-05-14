package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// decodes a JSON request body into a given struct
func DecodeJsonBody[T any](w http.ResponseWriter, r *http.Request, v *T) (ok bool) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(v); err != nil {
		log.Printf("Error decoding parameters: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	return true
}

// handles HTTP error response with an appropriate message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	RespondWithJson(w, code, map[string]string{"error": msg})
}

// sends a JSON response with the given status code and payload
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(resp)
}
