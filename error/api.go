package error

import (
	"net/http"
)

const (
	CodeInternalServerError = 1
	CodeValidationError     = 2
	CodeBadRequestError     = 3

	NameInternalServerError = "InternalServerError"
	NameValidationError     = "ValidationError"
	NameBadRequestError     = "BadRequestError"
	NameZeroNumberError     = "ZeroNumberError"
)

type ApiError struct {
	Message    string
	Name       string
	Code       int
	StatusCode int
	BaseError  error
}

func NewBadRequestError(msg string) ApiError {
	return ApiError{
		Message:    msg,
		Name:       NameBadRequestError,
		Code:       CodeBadRequestError,
		StatusCode: http.StatusBadRequest,
	}
}
