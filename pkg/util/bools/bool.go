package bools

import "github.com/herfurthe/golang.utils/pkg/util/slices"

func Not[Bool ~bool](value Bool) Bool {
	return !value
}

func And[Bool ~bool](values ...Bool) bool {
	if slices.IsEmpty(values) {
		return false
	}

	for _, value := range values {
		if Not(value) {
			return false
		}
	}

	return true
}

func Or[Bool ~bool](values ...Bool) bool {
	if slices.IsEmpty(values) {
		return false
	}

	for _, value := range values {
		if value {
			return true
		}
	}

	return false
}
