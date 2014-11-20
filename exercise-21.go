package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Printf("math.Sqrt(0) -> %v\n", math.Sqrt(0))
	fmt.Printf("math.Sqrt(2) -> %v\n", math.Sqrt(2))
	fmt.Printf("math.Sqrt(-4) -> %v\n", math.Sqrt(-4))
	fmt.Printf("sqrt(2) -> %v\n", sqrt(2))
	fmt.Printf("sqrt(-4) -> %v\n", sqrt(-4))
}
