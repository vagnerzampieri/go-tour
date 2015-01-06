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

func accumulate(operations []Operation) (accumulator int) {
	for _, op := range operations {
		value := op.Calc()
		fmt.Printf("%v %d\n", op, value)
		accumulator += value
	}

	return
}

func main() {
	var sum Operation
	fmt.Println("sum ->", sum)

	sum = Sum{10, 20}
	fmt.Println("sum = Sum{10, 20} ->", sum)

	fmt.Printf("%v %d\n", sum, sum.Calc())

	operations := make([]Operation, 4)
	operations[0] = Sum{30, 40}
	operations[1] = Subtraction{30, 15}
	operations[2] = Subtraction{10, 50}
	operations[3] = Sum{5, 2}

	fmt.Println("Valor acumulado =", accumulate(operations))
}
