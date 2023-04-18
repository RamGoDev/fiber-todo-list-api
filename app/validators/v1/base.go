package validators_v1

import (
	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()
