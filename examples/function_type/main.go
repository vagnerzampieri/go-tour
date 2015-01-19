package main

import "fmt"

type Aggregator func(n, m int) int

func Aggregate(values []int, init int, fn Aggregator) int {
	aggregated := init

	for _, v := range values {
		aggregated = fn(v, aggregated)
	}

	return aggregated
}

func CalcSum(values []int) int {
	sum := func(n, m int) int {
		return n + m
	}

	return Aggregate(values, 0, sum)
}

func CalcProduct(values []int) int {
	mult := func(n, m int) int {
		return n * m
	}

	return Aggregate(values, 1, mult)
}

func main() {
	values := []int{3, -2, 5, 7, 8, 22, 32, -1}

	fmt.Println(CalcSum(values))
	fmt.Println(CalcProduct(values))
}
