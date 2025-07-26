package set

import "github.com/herfurthe/golang.utils/pkg/util/bools"

type Set[Item comparable] map[Item]bool

func New[Item comparable](items ...Item) Set[Item] {
	set := make(Set[Item])
	set.InsertAll(items...)
	return set
}

func (s Set[Item]) Len() int {
	return len(s)
}

func (s Set[Item]) IsEmpty() bool {
	return s.Len() <= 0
}

func (s Set[Item]) Exists(item Item) bool {
	_, ok := s[item]
	return ok
}

func (s Set[Item]) ExistsAll(items ...Item) bool {
	for _, item := range items {
		if bools.Not(s.Exists(item)) {
			return false
		}
	}
	return true
}

func (s Set[Item]) ExistsAny(items ...Item) bool {
	for _, item := range items {
		if s.Exists(item) {
			return true
		}
	}
	return false
}

func (s Set[Item]) Insert(item Item) {
	s[item] = true
}

func (s Set[Item]) InsertAll(items ...Item) {
	for _, item := range items {
		s.Insert(item)
	}
}

func (s Set[Item]) Delete(item Item) {
	delete(s, item)
}

func (s Set[Item]) DeleteAll(items ...Item) {
	for _, item := range items {
		s.Delete(item)
	}
}

func (s Set[Item]) AsSlice() []Item {
	slices := make([]Item, 0, s.Len())
	for item := range s {
		slices = append(slices, item)
	}
	return slices
}
