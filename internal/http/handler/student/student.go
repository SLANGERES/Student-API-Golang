package student

import (
	"encoding/json"
	"errors"
	"fmt"
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

		// Handle missing body
		if errors.Is(err, io.EOF) {
			util.WriteJson(w, http.StatusBadRequest, util.GeneralErrorResponse(err))
			return
		}

		// Handle other decode errors
		if err != nil {
			util.WriteJson(w, http.StatusBadRequest, util.GeneralErrorResponse(err))
			return
		}

		// ✅ Validate request
		err = validator.New().Struct(std)
		if err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				fmt.Println(validationErrors)
				util.WriteJson(w, http.StatusBadRequest, util.ValidationError(validationErrors))
			} else {
				// Optional: catch other validation-related errors
				util.WriteJson(w, http.StatusInternalServerError, util.GeneralErrorResponse(err))
			}
			return
		}

		// 🎉 Success
		util.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
