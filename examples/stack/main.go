package main

import (
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Create stack with length", stack.Length())
	fmt.Println("Empty?", stack.Empty())
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Length() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Length() == 0
}
