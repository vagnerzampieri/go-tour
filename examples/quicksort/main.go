package main

import (
	"flag"
	"fmt"
	"os"
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

func main() {
	fmt.Println("numbers", numbers)
	fmt.Println("len numbers", lenNumbers)
	fmt.Println("numbers[0]", numbers[0])

	// tenho que dividir o slice, e pegar o valor do meio, tem que tomar cuidado quando for impar e par
	indexNumber := (lenNumbers / 2)
	chosen := numbers[indexNumber]
	newSlice := []string{}

	fmt.Println("indexNumber", indexNumber)
	fmt.Println("chosen", chosen)

	newSlice = append(numbers[:indexNumber], numbers[indexNumber+1:]...)

	fmt.Println("newSlice", newSlice)

	for i, n := range newSlice {
		fmt.Println(i, n)
	}
}
