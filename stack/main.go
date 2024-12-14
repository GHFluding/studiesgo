package main

import (
	"errors"
	"fmt"
)

type Stack []int

// добавление в стек
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// удаление верхнего элемента в стеке
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Стек пустой")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

// возращение верхнего элемента стека
func (s *Stack) Top() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("Cтек пустой")
	}
	return (*s)[len(*s)-1], nil
}

// пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	var stack Stack
	stack.Push(7)
	stack.Push(16)
	stack.Push(23)
	stack.Push(19)

	peek, _ := stack.Top()

	fmt.Printf("Верхняя часть стека: %v\n", peek)
	fmt.Printf("Содержимое стека: %v\n", stack)

	popped, _ := stack.Pop()
	fmt.Printf("Верхний элемент стека: %v\n", popped)
	fmt.Printf("Стек после применения метода Pop(): %+v", stack)
}
