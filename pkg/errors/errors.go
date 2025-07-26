package errors

import (
	"errors"
	"fmt"
)

func New(msg string) error {
	return errors.New(msg)
}

func Is(src, tar error) bool {
	return errors.Is(src, tar)
}

func Join(err ...error) error {
	return errors.Join(err...)
}

func Wrap(msg string, err error) error {
	return fmt.Errorf("%w: %s", err, msg)
}

func WrapF(format string, err error, msg ...any) error {
	return Wrap(fmt.Sprintf(format, msg...), err)
}
