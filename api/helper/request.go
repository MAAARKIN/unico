package helper

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

func BindJson(r *http.Request, destination interface{}) error {
	e := render.DecodeJSON(r.Body, destination)
	if e != nil {
		return ErrDecodeJson
	}
	return validateJsonFields(destination)
}

func validateJsonFields(input interface{}) error {

	validator := validator.New()

	if err := validator.Struct(input); err != nil {
		return err
	}

	return nil
}
