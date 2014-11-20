package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Small -> %v\n", Small)
	fmt.Printf("Small needInt -> %v\n", needInt(Small))
	fmt.Printf("Small needFloat -> %v\n", needFloat(Small))
	fmt.Printf("Small needFloat -> %v\n", needFloat(Big))
}
