package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var customMessages = map[string]string{
	"required": "/f is required",
	"email":    "/f must be a valid email address",
	"username": "/f must be a valid username",
	"numeric":  "/f must be a number",
	"alphanum": "/f must contain alphanumeric characters",
	"boolean":  "/f must be a boolean value",
	"uuid":     "/f must be a valid UUID",
	"min":      "/f must have at least /p characters",
	"max":      "/f must not exceed /p characters",
	"gte":      "/f must be at least /p",
	"lte":      "/f must be at most /p",
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	err := v.RegisterValidation("username", usernameValidation)
	PanicIfErr(err)

	return &Validator{validator: v}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)

	if err != nil {
		return NewHttpError(http.StatusBadRequest, "Bad Request", err)
	}

	return nil
}

func usernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	regex := regexp.MustCompile("^[a-zA-Z0-9._]+$")

	return regex.MatchString(username)
}

func FormatValidationErrors(errors validator.ValidationErrors) map[string]string {
	formattedErrors := make(map[string]string)

	for _, err := range errors {
		field := LcFirst(err.Field())
		tag := err.Tag()
		param := err.Param()

		msgTemplate := customMessages[tag]

		if msgTemplate == "" {
			msgTemplate = fmt.Sprintf("%s validation failed on %s", field, tag)
		}

		msg := strings.Replace(msgTemplate, "/f", field, -1)
		msg = strings.Replace(msg, "/p", param, -1)

		formattedErrors[field] = msg
	}

	return formattedErrors
}
