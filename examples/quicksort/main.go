package main

import (
	"flag"
	"fmt"
	"os"
)

var numbers []string

func init() {
	flag.Int("numbers", 10, "Add numbers to be sorted")
	flag.Parse()
}

func main() {
	numbers = os.Args
	fmt.Println("numbers", numbers)
}
