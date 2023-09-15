package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

func main() {
	// Create a stack for integers
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	// Create a stack for strings
	stringStack := NewStack[string]()
	stringStack.Push("apple")
	stringStack.Push("banana")
	stringStack.Push("cherry")

	// Pop and print items from the integer stack
	for {
		item, err := intStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	// Pop and print items from the string stack
	for {
		item, err := stringStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
