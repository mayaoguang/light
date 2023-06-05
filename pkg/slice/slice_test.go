package slice

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	s := NewLockSlice[string](10)
	s.Append("a")
	s.Append("b")
	fmt.Println(s.Len())
	fmt.Println(s.array)
	err := s.ProcessAndClear(Pro)
	fmt.Println(err)
	fmt.Println(s.Len())
	fmt.Println(s.array)
}

func Pro(s []string) error {
	fmt.Println("------", s)
	return nil
}
