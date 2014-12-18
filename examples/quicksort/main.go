package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	numbers    []string
	lenNumbers int
)

func init() {
	flag.Int("numbers", 10, "Add numbers to be sorted")
	flag.Parse()

	numbers = flag.Args()
	lenNumbers = len(numbers)

	errorArgs()
}

func errorArgs() {
	if lenNumbers == 0 {
		fmt.Println("Need numbers as arguments.\nEx.: go run main.go 1 2 3")
		os.Exit(1)
	} else if lenNumbers == 1 {
		fmt.Println("Returns himself", numbers)
		os.Exit(1)
	}
}

func errorOnlyNumbers(err error) {
	if err != nil {
		fmt.Println("Only numbers")
		os.Exit(1)
	}
}

func removeMidleValue() (chosen string, newSlice []string) {
	indexNumber := (lenNumbers / 2)
	chosen = numbers[indexNumber]
	newSlice = append(numbers[:indexNumber], numbers[indexNumber+1:]...)
	return
}

func partition(chosen string, slice []string) (lower, higher []int) {
	c, err := strconv.Atoi(chosen)
	errorOnlyNumbers(err)

	for _, n := range slice {
		cho, errors := strconv.Atoi(n)
		errorOnlyNumbers(errors)

		if cho < c {
			lower = append(lower, cho)
		} else if cho >= c {
			higher = append(higher, cho)
		}
	}
	return
}

func main() {
	fmt.Println("numbers", numbers)
	fmt.Println("len numbers", lenNumbers)
	fmt.Println("numbers[0]", numbers[0])

	chosen, newSlice := removeMidleValue()
	fmt.Println("chosen", chosen)
	fmt.Println("newSlice", newSlice)

	lower, higher := partition(chosen, newSlice)
	fmt.Println("lower", lower)
	fmt.Println("higher", higher)

	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
