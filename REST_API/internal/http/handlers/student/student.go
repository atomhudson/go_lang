package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/atomhudson/go_lang/student-api/internal/types"
	"github.com/atomhudson/go_lang/student-api/internal/utils/response"
	"github.com/go-playground/validator"
)

func NewStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(`router.HandleFunc("POST /api/students", student.NewStudent())`)
		slog.Info("creating a student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			slog.Info("empty request body")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, "empty request body"))
			return
		}
		if err != nil {
			slog.Info("invalid request body")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, "invalid request body"))
			return
		}

		// validate request
		if err := validator.New().Struct(student); err != nil {
			slog.Info("validation error")
			validatorErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatorErrs))
			return
		}

		slog.Info("student created")
		response.WriteJson(w, http.StatusCreated, response.Success(map[string]string{"message": "student created"}))
	}
}
