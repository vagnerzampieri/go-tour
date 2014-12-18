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
	fmt.Println("numbers", numbers)
}
