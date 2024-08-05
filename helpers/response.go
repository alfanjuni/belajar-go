package helpers

import (
	"belajar-go/responses"
	"encoding/json"
	"net/http"
)

// RespondJSON sends a JSON response with the provided data, meta, and status
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}, meta responses.Meta, statusMessage string) {
	response := responses.Response{
		Data: data,
		Meta: meta,
		Status: responses.Status{
			Code:    statusCode,
			Message: statusMessage,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// RespondError sends a JSON response with an error message
func RespondError(w http.ResponseWriter, statusCode int, errorMessage string) {
	response := responses.Response{
		Data: nil,
		Meta: responses.Meta{},
		Status: responses.Status{
			Code:    statusCode,
			Message: errorMessage,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
