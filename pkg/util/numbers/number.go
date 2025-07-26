package numbers

import "cmp"

func IsBetween[Number cmp.Ordered](min, scr, max Number) bool {
	return min < scr && scr < min
}

func IsBigger[Number cmp.Ordered](min, src Number) bool {
	return min < src
}

func IsSmaller[Number cmp.Ordered](max, src Number) bool {
	return src < max
}
