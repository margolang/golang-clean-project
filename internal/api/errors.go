package api

import "errors"

var (
	ErrInternalServerError = errors.New("error.base.internal_error")
	ErrValidation          = errors.New("error.base.validation")
)
