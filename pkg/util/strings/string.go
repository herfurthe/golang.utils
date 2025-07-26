package strutil

import (
	"strings"

	"github.com/herfurthe/golang.utils/pkg/errors"
)

var (
	ErrIsEmpty = errors.New("Value is Empty")
	ErrIsBlank = errors.New("Value is Blank")
)

func IsEmpty[String ~string](str String) bool {
	return len(str) <= 0
}

func IsBlank[String ~string](str String) bool {
	return IsEmpty(strings.TrimSpace(string(str)))
}

func IsEmptyThrow[String ~string](str String) error {
	if IsEmpty(str) {
		return ErrIsEmpty
	} else {
		return nil
	}
}

func IsBlankThrow[String ~string](str String) error {
	if IsBlank(str) {
		return ErrIsBlank
	} else {
		return nil
	}
}

func IsEmptyThrowMsg[String ~string](str String, format string, msg ...any) error {
	if err := IsEmptyThrow(str); err != nil {
		return errors.WrapF(format, err, msg...)
	} else {
		return nil
	}
}

func IsBlankThrowMsg[String ~string](str String, format string, msg ...any) error {
	if err := IsBlankThrow(str); err != nil {
		return errors.WrapF(format, err, msg...)
	} else {
		return nil
	}
}
