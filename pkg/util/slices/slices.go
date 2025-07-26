package slices

func IsEmpty[Item any](items []Item) bool {
	return len(items) <= 0
}
