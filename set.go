package main

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](initialElements ...T) Set[T] {
	s := make(Set[T], len(initialElements))
	s.Insert(initialElements...)
	return s
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Insert(t ...T) {
	for _, t := range t {
		s[t] = struct{}{}
	}
}

func (s Set[T]) Remove(t ...T) {
	for _, t := range t {
		delete(s, t)
	}
}
