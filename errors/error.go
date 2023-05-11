package errors

import "errors"

var (
	ErrInvalid   = errors.New("invalid")
	ErrNotFound  = errors.New("not found")
	ErrExist     = errors.New("exist")
	ErrForbidden = errors.New("forbidden")
	ErrInternal  = errors.New("internal")
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
