package httpError

import (
	"backend/service"
	"errors"
	"net/http"
)

type APIError struct {
	Status  int
	Message string
}

func FromError(err error) APIError {
	var apiErr APIError
	var svcError service.Error
	if errors.As(err, &svcError) {
		apiErr.Message = svcError.AppError().Error()
		svcErr := svcError.SvcError()
		switch svcErr {
		case service.ErrBadRequest:
			apiErr.Status = http.StatusBadRequest
		case service.ErrInternalFailure:
			apiErr.Status = http.StatusInternalServerError
		case service.ValidateErr:
			apiErr.Status = http.StatusForbidden
		case service.ErrNotfound:
			apiErr.Status = http.StatusNotFound
		}
	}
	return apiErr
}
