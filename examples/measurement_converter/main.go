// Use:
// go run main.go -h -> help
// go run main.go 10.08 celsius
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/vagnerzampieri/go-tour/examples/measurement_converter/converter"
)

var (
	lenArgs           int
	originValues      []string
	originMeasurement string
)

func init() {
	flag.Float64("value", 10.07, "Adding value to converter")
	flag.String("measurement", "celsius or fahrenheit", "Adding conversor")
	flag.Parse()

	lenArgs = len(flag.Args())
	lastArg := len(flag.Args()) - 1

	validArgs()

	originMeasurement = flag.Args()[lastArg]
	originValues = flag.Args()[0:lastArg]
}

func validArgs() {
	if lenArgs < 2 {
		fmt.Println("Use: converter <value> <measurement>")
		os.Exit(1)
	}
}

func destinyMeasurement() (str string) {
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

func destinyValue(originValue float64) (destinyValue float64) {
	if originMeasurement == "celsius" {
		destinyValue = converter.Celsius(originValue)
	} else {
		destinyValue = converter.Kilometre(originValue)
	}
	return
}

func main() {
	for i, v := range originValues {
		originValue, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("The value %s in position %d is not a valid number!\n", v, i)
			os.Exit(1)
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			originValue, originMeasurement, destinyValue(originValue), destinyMeasurement())
	}
}
