package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	value       *float64
	measurement *string
)

func init() {
	value = flag.Float64("value", 10.07, "Adding value to converter")
	measurement = flag.String("measurement", "celsius or fahrenheit", "Adding conversor")
	flag.Parse()
}

func main() {
	//fmt.Println("len(os.Args) ->", len(os.Args))
	//fmt.Println("os.Args ->", os.Args)
	//fmt.Println("os.Args[0] ->", os.Args[0])

	if len(os.Args) < 3 {
		fmt.Println("Use: converter <value> <measurement>")
		os.Exit(1)
	}

	originMeasurement := os.Args[len(os.Args)-1]
	originValue := os.Args[1 : len(os.Args)-1]

	fmt.Println("originMeasurement := os.Args[len(os.Args)-1]", originMeasurement)
	fmt.Println("originValue := os.Args[1 : len(os.Args)-1]", originValue)
}
