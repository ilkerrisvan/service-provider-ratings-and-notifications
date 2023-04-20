package validator

import (
	"github.com/go-playground/validator/v10"
)

type StructLevelType validator.StructLevel
type IValidator interface {
	ValidateStruct(s interface{}) error
	RegisterStructLevelValidation(f func(sl StructLevelType), m interface{})
}

type Validator struct {
	validate *validator.Validate
}

func New() IValidator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}

func (v *Validator) RegisterStructLevelValidation(f func(sl StructLevelType), m interface{}) {
	v.validate.RegisterStructValidation(func(sl validator.StructLevel) {
		f(sl)
	}, m)
}
