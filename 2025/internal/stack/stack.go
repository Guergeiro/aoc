package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	Print()
	PopAll() []T
}

type stack[T any] struct {
	data []T
}

func NewStack[T any](initial []T) Stack[T] {
	return &stack[T]{
		data: initial,
	}
}

func (s *stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) Pop() (T, error) {
	if s.Size() > 0 {
		value := s.data[s.Size()-1]
		s.data = s.data[:s.Size()-1]
		return value, nil
	}

	var zero T

	return zero, errors.New("no such element")
}

func (s *stack[T]) Peek() (T, error) {
	if s.Size() > 0 {
		value := s.data[s.Size()-1]
		return value, nil
	}

	var zero T

	return zero, errors.New("no such element")
}

func (s *stack[T]) Size() int {
	return len(s.data)
}

func (s *stack[T]) Print() {
	fmt.Println(s.data)
}

func (s *stack[T]) PopAll() []T {
	output := []T{}
	for s.Size() > 0 {
		if val, err := s.Pop(); err == nil {
			output = append(output, val)
		}
	}
	return output
}
