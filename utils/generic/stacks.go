package generic

import "fmt"

// 泛型 链表实现栈
type (
	Stacks[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		data T
		prev *node[T]
	}
)

func New[T any]() *Stacks[T] {
	return &Stacks[T]{nil, 0}
}

func (slf *Stacks[T]) Len() int {
	return slf.length
}

func (slf *Stacks[T]) IsEmpty() bool {
	return slf.length == 0
}

func (slf *Stacks[T]) Peek() (t T, err error) {
	if slf.length == 0 {
		return t, fmt.Errorf("stack is empty")
	}
	return slf.top.data, nil
}

func (slf *Stacks[T]) Pop() (t T, err error) {
	if slf.length == 0 {
		return t, fmt.Errorf("stacks is empty")
	}
	v := slf.top
	slf.top = v.prev
	slf.length--
	return v.data, nil
}

func (slf *Stacks[T]) Push(data T) {
	newNode := &node[T]{data: data, prev: slf.top}
	slf.top = newNode
	slf.length++
}
