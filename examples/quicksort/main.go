package main

import (
	"flag"
)

var numbers *int

func init() {
	numbers = flag.Int("numbers", 10, "Adding integers to create slice")
	flag.Parse()
}

func main() {}
