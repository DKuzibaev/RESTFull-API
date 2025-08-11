package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValidate[T any](payload T) error {
	// проверка email
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
