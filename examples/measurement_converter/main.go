package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use: converter <value> <measurement>")
		os.Exit(1)
	}
}
