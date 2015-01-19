package main

import "fmt"

func Aggregate(values []int, init int, fn func(n, m int) int) int {
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

func main() {

}
