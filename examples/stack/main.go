package main

import (
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Create stack with length", stack.Length())
	fmt.Println("Empty?", stack.Empty())

	stack.Add("Go")
	stack.Add(2009)
	stack.Add(3.14)
	stack.Add("End")
	fmt.Println("Length after 4 values:", stack.Length())
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

func (stack *Stack) Add(value interface{}) {
	stack.values = append(stack.values, value)
}
