package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)

	if err != nil {
		var message []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				message = append(message, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				message = append(message, fmt.Sprintf("%s is not valid email",
					err.Field()))
			case "min":
				message = append(message, fmt.Sprintf("%s character must be %s digits in minimal",
					err.Field(), err.Param()))
			}
		}

		err := errors.New(strings.Join(message, ","))
		return err
	}

	return nil
}

func (v Validator) NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}
