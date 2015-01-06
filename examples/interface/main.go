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
	return fmt.Sprintf("%d + %d ->", s.value1, s.value2)
}

func main() {

}
