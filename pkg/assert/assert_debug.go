//go:build debug

package assert

import (
	"fmt"

	"github.com/herfurthe/golang.utils/pkg/errors"
	"github.com/herfurthe/golang.utils/pkg/util/bools"
)

func ThatTrue(mustBeTrue bool, format string, msg ...any) {
	if bools.Not(mustBeTrue) {
		panic(fmt.Sprintf(format, msg...))
	}
}

func ThatTrueLazy(mustBeTrue func() bool, format string, msg ...any) {
	ThatTrue(mustBeTrue(), format, msg...)
}

func ThatFalse(mustBeFalse bool, format string, msg ...any) {
	if mustBeFalse {
		panic(fmt.Sprintf(format, msg...))
	}
}

func ThatFalseLazy(mustBeFalse func() bool, format string, msg ...any) {
	ThatFalse(mustBeFalse(), format, msg...)
}

func WithoutError(msg string, errs ...error) {
	if joinedErr := errors.Join(errs...); joinedErr != nil {
		panic(fmt.Sprintf("%s %s", joinedErr.Error(), msg))
	}
}

func WithoutErrorLazy(errors func() error, msg string) {
	WithoutError(msg, errors())
}
