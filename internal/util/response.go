package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/go-playground/validator"
	"github.com/slangeres/Student-api-go/internal/types"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

// These are genrally a wrapper that wrap the reponse
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Error wrapper
func GeneralErrorResponse(err error) types.ErrorResponse {
	var response types.ErrorResponse

	response.Status = StatusError
	response.Error = err.Error()

	return response
}

// Validating error response
func ValidationError(errs validator.ValidationErrors) types.ErrorResponse {
	var errMsg []string

	
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s is reqired", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("%s is required ", err.Field()))
		}
	}
	return types.ErrorResponse{
		Status: "OK",
		Error:  strings.Join(errMsg, ","),
	}
}
