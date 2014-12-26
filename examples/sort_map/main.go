package main

import (
	"fmt"
	"sort"
)

func main() {
	squares := make(map[int]int, 15)

	for n := 1; n <= 15; n++ {
		squares[n] = n * n
	}

	fmt.Println("squares ->", squares)

	numbers := make([]int, 0, len(squares))

	for n := range squares {
		numbers = append(numbers, n)
	}

	fmt.Println("numbers ->", numbers)

	sort.Ints(numbers)

	fmt.Println("sorted numbers ->", numbers)

	for _, number := range numbers {
		square := squares[number]
		fmt.Printf("%d^2 = %d\n", number, square)
	}
}
