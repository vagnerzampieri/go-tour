package main

import "fmt"

type Operation interface {
	Calc() int
}

type Sum struct {
	value1, value2 int
}

func (s Sum) Calc() int {
	return s.value1 + s.value2
}

func (s Sum) String() string {
	return fmt.Sprintf("%d + %d =", s.value1, s.value2)
}

type Subtraction struct {
	value1, value2 int
}

func (s Subtraction) Calc() int {
	return s.value1 - s.value2
}

func (s Subtraction) String() string {
	return fmt.Sprintf("%d - %d =", s.value1, s.value2)
}

func main() {
	var sum Operation
	fmt.Println("sum ->", sum)

	sum = Sum{10, 20}
	fmt.Println("sum = Sum{10, 20} ->", sum)

	fmt.Printf("%v %d\n", sum, sum.Calc())
}
