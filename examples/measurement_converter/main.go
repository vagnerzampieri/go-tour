package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	value             *float64
	measurement       *string
	lenArgs           int
	originValue       []string
	originMeasurement string
)

func init() {
	value = flag.Float64("value", 10.07, "Adding value to converter")
	measurement = flag.String("measurement", "celsius or fahrenheit", "Adding conversor")
	flag.Parse()

	validArgs()

	originMeasurement = os.Args[lenArgs-1]
	originValue = os.Args[1 : lenArgs-1]
}

func validArgs() {
	lenArgs = len(os.Args)

	if lenArgs < 3 {
		fmt.Println("Use: converter <value> <measurement>")
		os.Exit(1)
	}
}

func destinyMeasurement() (str string) {
	fmt.Println("originMeasurement := os.Args[len(os.Args)-1]", originMeasurement)
	fmt.Println("originValue := os.Args[1 : len(os.Args)-1]", originValue)

	if originMeasurement == "celsius" {
		str = "fahrenheit"
	} else if originMeasurement == "kilometre" {
		str = "miles"
	} else {
		str = originMeasurement
		fmt.Printf("%s is not a unit of measurement known!\n", str)
		os.Exit(1)
	}
	return
}

func main() {
	//fmt.Println("len(os.Args) ->", len(os.Args))
	//fmt.Println("os.Args ->", os.Args)
	//fmt.Println("os.Args[0] ->", os.Args[0])
	destinyMeasurement()
}
