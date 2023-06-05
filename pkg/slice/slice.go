package slice

import "sync"

type LockSlice[T comparable] struct {
	array []T
	lock  sync.Mutex
}

func NewLockSlice[T comparable](l int) *LockSlice[T] {
	return &LockSlice[T]{
		array: make([]T, 0, l),
		lock:  sync.Mutex{},
	}
}

func (s *LockSlice[T]) Len() int {
	return len(s.array)
}

func (s *LockSlice[T]) Clear() *LockSlice[T] {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = s.array[0:0]
	return s
}

func (s *LockSlice[T]) Append(v T) *LockSlice[T] {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = append(s.array, v)
	return s
}

func (s *LockSlice[T]) ProcessAndClear(f func([]T) error) (*LockSlice[T], error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if err := f(s.array); err != nil {
		return s, err
	}

	s.array = s.array[0:0]
	return s, nil
}
