package main

import (
	"errors"
	"fmt"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Стек пустой")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

func (s *Stack) Peek() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Cтек пустой")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	var stack Stack
	stack.Push(1)
	stack.Push(10)
	stack.Push(22)
	stack.Push(18)

	peek, _ := stack.Peek()

	fmt.Printf("Верхняя часть стека: %v\n", peek)
	fmt.Printf("Содержимое стека: %v\n", stack)

	popped, _ := stack.Pop()
	fmt.Printf("Верхний элемент стека: %v\n", popped)
	fmt.Printf("Стек после применения метода Pop(): %+v", stack)
}
