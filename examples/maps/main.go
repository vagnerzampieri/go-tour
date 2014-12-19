// Use:
// go run main.go -h
// go run main.go John Doe and Marie Doe
package main

import (
	"flag"
	"fmt"
)

var words []string

func init() {
	flag.String("words", "John Doe and Marie Doe", "Add only string")
	flag.Parse()

	words = flag.Args()
}

func main() {
	fmt.Println("words ->", words)
}
