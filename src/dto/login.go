package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (input LoginInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Password, validation.Required),
		validation.Field(&input.Email, validation.Required, is.Email),
	)
}
