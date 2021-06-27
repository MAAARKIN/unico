package helper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
)

const maxStackTraceSize = 5

var errorHandlerMap = make(map[error]int)

var ErrDecodeJson = errors.New("error decoding json")
var ErrMarshalJson = errors.New("error generating json")
var ErrJsonValidation = errors.New("error validating json")

func init() {
	errorHandlerMap[db.ErrRecordNotFound] = http.StatusBadRequest
	errorHandlerMap[ErrDecodeJson] = http.StatusBadRequest
	errorHandlerMap[ErrMarshalJson] = http.StatusBadRequest
	errorHandlerMap[domain.ErrFeiraWithRegistroAlreadyExist] = http.StatusBadRequest

	//register how to deal with errors here to HTTP layer
}

type HttpError struct {
	Status      int      `json:"-"`
	Description string   `json:"description,omitempty"`
	Messages    []string `json:"messages,omitempty"`
}

type StackTracer interface {
	StackTrace() errors.StackTrace
}

func DealWith(err error) HttpError {
	if ok, httpError := IsValidationError(err); ok {
		return *httpError
	} else if errCode, ok := errorHandlerMap[err]; ok {
		return HttpError{Status: errCode, Description: err.Error()}
	} else {
		if errWithStack, ok := err.(StackTracer); ok {
			size := len(errWithStack.StackTrace())
			if size > maxStackTraceSize {
				size = maxStackTraceSize
			}
			result := fmt.Sprintf("%+v", errWithStack.StackTrace()[0:size])

			log.Printf(
				`[ERROR] Internal error, check stacktrace to see the problem. 
				Cause: %v - Stacktrace: %s`,
				err,
				result,
			)
		} else {
			log.Printf("[ERROR] Internal error, check stacktrace to see the problem: %v", err)
		}

		return HttpError{
			Status:      http.StatusInternalServerError,
			Description: "Internal error, please report to unico Team",
		}
	}
}

func IsValidationError(err error) (bool, *HttpError) {
	v := &validator.ValidationErrors{}
	if errors.As(err, v) {
		validationErrors := HttpError{Status: http.StatusBadRequest, Description: ErrJsonValidation.Error()}
		for _, err := range err.(validator.ValidationErrors) {
			message := generateErrorMessage(err)
			validationErrors.Messages = append(validationErrors.Messages, message.Error())
		}
		return true, &validationErrors
	}
	return false, nil
}

func generateErrorMessage(err validator.FieldError) error {
	return fmt.Errorf("error: field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
}
