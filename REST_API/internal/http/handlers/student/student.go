package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/atomhudson/go_lang/student-api/internal/storage"
	"github.com/atomhudson/go_lang/student-api/internal/types"
	"github.com/atomhudson/go_lang/student-api/internal/utils/response"
	"github.com/go-playground/validator"
)

func NewStudent(storage storage.Storage) http.HandlerFunc {
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

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Phone, student.Age)
		if err != nil {
			slog.Info("failed to create student")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, "failed to create student"))
			return
		}

		slog.Info("student created with id", slog.Int("id", lastId))
		response.WriteJson(w, http.StatusCreated, response.Success(map[string]string{"userId": strconv.Itoa(lastId)}))
	}
}

func GetStudents(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting student by id", slog.String("id", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("invalid id")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, "invalid id"))
			return
		}
		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Info("student not found")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, "student not found"))
			return
		}
		slog.Info("student found", slog.Any("student", student))
		response.WriteJson(w, http.StatusOK, response.Success(student))
	}
}

func GetStudentsList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(`router.HandleFunc("GET /api/students", student.GetStudents())`)
		slog.Info("getting all students")
		students, err := storage.GetStudentsList()
		if err != nil {
			slog.Info("failed to get students")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, "failed to get students"))
			return
		}
		slog.Info("students found", slog.Any("students", students))
		response.WriteJson(w, http.StatusOK, response.Success(students))
	}
}
