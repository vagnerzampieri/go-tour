package main

import (
	"fmt"
	"time"
)

func GenerateFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b
			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}

func main() {
}
