package util

import (
	"encoding/json"
	"net/http"

	"github.com/slangeres/Student-api-go/internal/types"
)

const (
	StatusOK="OK"
	StatusError="Error"
)

//These are genrally a wrapper that wrap the reponse
func WriteJson(w http.ResponseWriter,status int,data interface{}) error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
//Error wrapper
func GeneralErrorResponse(err error)types.ErrorResponse{
	var response types.ErrorResponse

	response.Status=StatusError
	response.Error=err.Error()

	return response
}