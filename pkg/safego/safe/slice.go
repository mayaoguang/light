package safe

import (
	"sync"
)

type SafeSlice struct {
	array []string
	lock  sync.Mutex
}

func NewSafeSlice(len int) *SafeSlice {
	return &SafeSlice{
		array: make([]string, 0, len),
		lock:  sync.Mutex{},
	}
}

func (s *SafeSlice) Append(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = append(s.array, v)
}

func (s *SafeSlice) Len() int {
	return len(s.array)
}

func (s *SafeSlice) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.array = s.array[0:0]
}

func (s *SafeSlice) ProcessAndClear(f func([]string) error, limit int) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Len() == 0 {
		return nil
	}
	if limit > 0 {
		for s.Len() > limit {
			if err := f(s.array[:limit]); err != nil {
				return err
			}
			s.array = s.array[limit:]
		}
	}
	if err := f(s.array); err != nil {
		return err
	}
	s.array = s.array[0:0]

	return nil
}
