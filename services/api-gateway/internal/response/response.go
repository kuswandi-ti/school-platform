package response

import (
	"encoding/json"
	"net/http"
)

type Envelope struct {
	Data  any        `json:"data"`
	Meta  any        `json:"meta"`
	Error *ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
}

const (
	CodeInternalError   = "INTERNAL_ERROR"
	CodeNotFound        = "NOT_FOUND"
	CodeValidationError = "VALIDATION_ERROR"
)

func JSON(w http.ResponseWriter, status int, data any, meta any) {
	write(w, status, Envelope{
		Data:  data,
		Meta:  meta,
		Error: nil,
	})
}

func Error(w http.ResponseWriter, status int, code string, message string, details any) {
	write(w, status, Envelope{
		Data: nil,
		Meta: nil,
		Error: &ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func write(w http.ResponseWriter, status int, payload Envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
