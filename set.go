package main

type Set[T comparable] map[T]struct{}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func (s Set[T]) Insert(t T) {
	s[t] = struct{}{}
}
