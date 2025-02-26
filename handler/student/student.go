package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Babur171/carZone-golang/models"
	"github.com/Babur171/carZone-golang/service"
	"github.com/Babur171/carZone-golang/store"
	"github.com/go-playground/validator/v10"
)

func New(storage store.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student models.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			service.WriteJson(w, http.StatusBadRequest, service.GeneralError(err))
			return
		}
		if err != nil {
			service.WriteJson(w, http.StatusBadRequest, service.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			errorType := err.(validator.ValidationErrors)

			service.WriteJson(w, http.StatusBadRequest, service.ValidtionError(errorType))
			return
		}
		lastId, err := storage.CreateStudent(
			student.Email,
			student.Name,
		)
		if err != nil {
			service.WriteJson(w, http.StatusInternalServerError, service.GeneralError(err))
			return
		}

		service.WriteJson(w, http.StatusCreated, map[string]interface{}{
			"success": "created new user",
			"student": lastId,
		})

	}
}
