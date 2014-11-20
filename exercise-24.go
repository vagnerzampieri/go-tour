package main

import (
	"fmt"
	"math"
)

func SqrtNewt(x float64) float64 {
	if x == 0 {
		return 0
	}

	z := float64(1)

	for i := 0; i < int(x); i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}

	return z
}

func main() {
	times := 15
	for i := 0; i < times; i++ {
		sqrt := math.Sqrt(float64(i))
		newt := SqrtNewt(float64(i))
		fmt.Println(i, "squared:")
		fmt.Println("  Sqrt:", sqrt)
		fmt.Println("  Newt:", newt)
		fmt.Println("  Difference:", math.Abs(sqrt-newt))
	}
}
