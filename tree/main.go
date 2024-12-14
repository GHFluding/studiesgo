package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ChekInp() (int, error) {
	var input string
	fmt.Scanln(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid input format")
	} else {
		return number, nil
	}
}

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}
	if t.val == value {
		return errors.New("This node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &Tree{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &Tree{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

// func (t *Tree) Sort( size int ){
// 	if t == nil {
// 		return errors.New("Tree is nil")
// 	}else{
// 	var count int
// 	var copy Tree
// 	for i :=1; i<size; i++{
// 		copy = *t
// 		count--
// 		for((count >= 0)&&(t.right.val > copy.val)){
// 			t.right.val = t.val
//             t.val = copy.val
//             count--
// 		}
// 	}
// }
// }

func (t *Tree) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

func (t *Tree) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

func (t *Tree) PrintTree() {
	if t == nil {
		return
	}
	t.left.PrintTree()
	fmt.Print(t.val)
	t.right.PrintTree()
}

func main() {
	var t Tree
	fmt.Println("Write the number of tree elements")
	number, err := ChekInp()
	if err != nil {
		fmt.Println(err)
		err = nil
	}
	fmt.Println("Write elements")
	for i := number; i != 0; i-- {
		var value int
		value, err = ChekInp()
		if err != nil {
			fmt.Println(err)
			err = nil
		}
		t.Insert(value)
	}
	fmt.Println("Tree:")
	t.PrintTree()
}
