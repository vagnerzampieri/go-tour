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
	fmt.Println("divisão", (lenNumbers / 2))
	fmt.Println("meio", numbers[(lenNumbers/2)])

	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
