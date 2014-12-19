package main

import (
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Create stack with length", stack.Length())
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Length() int {
	return len(stack.values)
}
