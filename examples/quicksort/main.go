package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var numbers []string

func init() {
	flag.Int("numbers", 10, "Add numbers to be sorted")
	flag.Parse()

	numbers = flag.Args()
	errorArgs(len(numbers))
}

func errorArgs(length int) {
	if length == 0 {
		fmt.Println("Need numbers as arguments.\nEx.: go run main.go 1 2 3")
		os.Exit(1)
	}
}

func errorOnlyNumbers(err error) {
	if err != nil {
		fmt.Println("Only numbers")
		os.Exit(1)
	}
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	nums := make([]int, len(numbers))
	copy(nums, numbers)

	indexNumber := (len(nums) / 2)
	chosen := nums[indexNumber]
	nums = append(nums[:indexNumber], nums[indexNumber+1:]...)

	lower, higher := partition(chosen, nums)

	return append(append(quicksort(lower), chosen), quicksort(higher)...)
}

func partition(chosen int, slice []int) (lower, higher []int) {
	for _, n := range slice {
		if n <= chosen {
			lower = append(lower, n)
		} else {
			higher = append(higher, n)
		}
	}
	return
}

func main() {
	nums := make([]int, len(numbers))

	for i, n := range numbers {
		number, err := strconv.Atoi(n)
		errorOnlyNumbers(err)
		nums[i] = number
	}
	fmt.Println(quicksort(nums))
}
