package main

import "fmt"

func Aggregate(values []int, init int, fn func(n, m int) int) {
	aggregated := init

	for _, v := range values {
		aggregated = fn(v, aggregated)
	}

	return aggregated
}

func main() {

}
