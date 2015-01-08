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

func Timer(f func()) {
	start := time.Now()

	f()

	fmt.Printf("\nTempo de execução: %s\n\n",
		time.Since(start))
}

func main() {
	Timer(GenerateFibonacci(8))
	Timer(GenerateFibonacci(48))
	Timer(GenerateFibonacci(88))
}
