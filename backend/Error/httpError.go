package httpError

import (
	"errors"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

func New(statusCode int, msg string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    msg,
	}
}

var (
	ErrBadRequest   = &APIError{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
	ErrUnauthorized = &APIError{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden    = &APIError{StatusCode: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound     = &APIError{StatusCode: http.StatusNotFound, Message: "Resource Not Found"}
	ErrInternal     = &APIError{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
	ErrInvalidInput = &APIError{StatusCode: http.StatusBadRequest, Message: "Invalid Input"}
	ErrLoginFailed  = &APIError{StatusCode: http.StatusUnauthorized, Message: "Email or Password is incorrect"}
)

func FromError(err error) *APIError {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr
	}
	return ErrInternal
}
