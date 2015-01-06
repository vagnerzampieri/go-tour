package main

import (
	"fmt"
	"time"
)

type Operation interface {
	Calc() int
}

type Age struct {
	birthdate int
}

func (a Age) Calc() int {
	return time.Now().Year() - a.birthdate
}

func (a Age) String() string {
	return fmt.Sprintf("Idade desde %d", a.birthdate)
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
	ages := make([]Operation, 3)
	ages[0] = Age{1969}
	ages[1] = Age{1977}
	ages[2] = Age{2001}

	fmt.Println("Idades acumuladas", accumulate(ages))
}
