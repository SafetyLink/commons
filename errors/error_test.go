package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	err := returnInvalidError()
	if errors.Is(err, ErrInvalid) {
		fmt.Print("invalid")
	} else {
		fmt.Print("not invalid")
	}

}

func returnInvalidError() error {
	return ErrInvalid
}

func returnError() error {

	return ErrNotFound
}
