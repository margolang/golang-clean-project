package validator

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationError struct {
	FieldError validator.FieldError
}

func (v *ValidationError) Tag() string {
	return v.FieldError.Tag()
}

func (v *ValidationError) ActualTag() string {
	return v.FieldError.ActualTag()
}

func (v *ValidationError) Namespace() string {
	return v.FieldError.Namespace()
}

func (v *ValidationError) StructNamespace() string {
	return v.FieldError.StructNamespace()
}

func (v *ValidationError) Field() string {
	return v.FieldError.Field()
}

func (v *ValidationError) StructField() string {
	return v.FieldError.StructField()
}

func (v *ValidationError) Value() interface{} {
	return v.FieldError.Value()
}

func (v *ValidationError) Param() string {
	return v.FieldError.Param()
}

func (v *ValidationError) Kind() reflect.Kind {
	return v.FieldError.Kind()
}

func (v *ValidationError) Type() reflect.Type {
	return v.FieldError.Type()
}

func (v *ValidationError) Translate(ut ut.Translator) string {
	return v.FieldError.Translate(ut)
}

func (v *ValidationError) Error() string {
	if prefix, ok := tagPrefixMap[v.Tag()]; ok {
		return prefix
	}
	return v.FieldError.Error()
}

func ToValidationContext(err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for i := 0; i < len(validationErrors); i++ {
			var fieldError validator.FieldError
			if errors.As(validationErrors[i], &fieldError) {
				validationErrors[i] = &ValidationError{
					FieldError: fieldError,
				}
			}
		}

		return validationErrors
	}

	return err
}
