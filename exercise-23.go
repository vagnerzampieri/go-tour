package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Printf("math.Pow(3, 2) -> %v\n", math.Pow(3, 2))
	fmt.Printf("math.Pow(3, 3) -> %v\n", math.Pow(3, 3))
	fmt.Printf("pow(3, 2, 10) -> %v\n", pow(3, 2, 10))
	fmt.Printf("pow(3, 3, 20) -> %v\n", pow(3, 3, 20))
}
