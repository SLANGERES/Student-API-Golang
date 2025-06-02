package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/slangeres/Student-api-go/internal/types"
	"github.com/slangeres/Student-api-go/internal/util"
)

func Home() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	}
}

func PostStudent() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var std types.Student

		slog.Info("Creating A new Student ........")

		err := json.NewDecoder(r.Body).Decode(&std)

		//this is specially for the missing values EOF
		if errors.Is(err, io.EOF) {
			util.WriteJson(w, http.StatusBadRequest, util.GeneralErrorResponse(err))

			return
		}

		if err!=nil{
			util.WriteJson(w,http.StatusBadRequest,util.GeneralErrorResponse(err))
		}

		// TODO Validate Request

		err=validator.New().Struct(std)
		if err!=nil{
			util.WriteJson(w,http.StatusBadRequest,util.GeneralErrorResponse(err))
		}

		util.WriteJson(w, http.StatusCreated, map[string]string{"sucess": "ok"})
	}
}
