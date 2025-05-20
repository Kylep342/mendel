package responses

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON function to respond with a JSON payload
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// RespondWithData is a wrapper function for responding on a successful request
func RespondWithData(w http.ResponseWriter, code int, data interface{}) {
	RespondWithJSON(w, code, map[string]interface{}{"data": data})
}

// RespondWithError is a wrapper function for responding on an unsuccessful request
func RespondWithError(w http.ResponseWriter, code int, data interface{}) {
	RespondWithJSON(w, code, map[string]interface{}{"error": data})
}
