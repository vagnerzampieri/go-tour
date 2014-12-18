package main

import (
	"flag"
	"fmt"
)

var numbers []string

func init() {
	flag.Int("numbers", 10, "Add numbers to be sorted")
	flag.Parse()
}

func main() {
	numbers = flag.Args()
	lenNumbers := len(numbers)

	fmt.Println("numbers", numbers)
	fmt.Println("len numbers", lenNumbers)
	fmt.Println("numbers[0]", numbers[0])

	// tenho que dividir o slice, e pegar o valor do meio, tem que tomar cuidado quando for impar e par

	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
