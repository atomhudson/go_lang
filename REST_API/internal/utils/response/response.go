package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/atomhudson/go_lang/student-api/internal/types"
	"github.com/go-playground/validator"
)

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, resp types.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func GeneralError(err error, data string) types.Response {
	return types.Response{
		Status:  StatusError,
		Message: err.Error(),
		Data:    data,
	}
}

func Success(data interface{}) types.Response {
	return types.Response{
		Status:  StatusOk,
		Message: "success",
		Data:    data,
	}
}

func ValidationError(errs validator.ValidationErrors) types.Response {
	var errMsg []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("field %s is required", err.Field()))

		default:
			errMsg = append(errMsg, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}
	return types.Response{
		Status:  StatusError,
		Message: "validation error",
		Data:    strings.Join(errMsg, ", "),
	}
}
