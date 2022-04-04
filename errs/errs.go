package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (this AppError) Error() string {
	return this.Message
}

func NewNotFoundError(message string) error {
	return AppError{Code: http.StatusNotFound, Message: message}
}

func NewUnexpectedError() error {
	return AppError{Code: http.StatusInternalServerError, Message: "Unexpected error"}
}

func NewValidationError(message string) error {
	return AppError{Code: http.StatusUnprocessableEntity, Message: message}
}
