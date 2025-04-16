package service

import "errors"

var (
	ErrBadRequest      = errors.New("Bad Requset")
	ErrInternalFailure = errors.New("Internal Failure")
	ValidateErr        = errors.New("Email or Password is incorrect")
	ErrNotfound        = errors.New("Not Found")
)

type Error struct {
	appError error
	svcError error
}

func NewError(svcErr, appErr error) error {
	return Error{
		svcError: svcErr,
		appError: appErr,
	}
}

func (e Error) AppError() error {
	return e.appError
}
func (e Error) SvcError() error {
	return e.svcError
}

func (e Error) Error() string {
	return errors.Join(e.appError, e.svcError).Error()
}
