//go:build !debug

package assert

func ThatTrue(_ bool, _ string, _ ...any) {
}

func ThatTrueLazy(_ func() bool, _ string, _ ...any) {
}

func ThatFalse(_ bool, _ string, _ ...any) {
}

func ThatFalseLazy(_ func() bool, _ string, _ ...any) {
}

func WithoutError(_ string, _ ...error) {
}

func WithoutErrorLazy(_ func() error, _ string) {}
