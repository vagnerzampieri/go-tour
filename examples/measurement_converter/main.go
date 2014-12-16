package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	value             *float64
	measurement       *string
	lenArgs           int
	originValues      []string
	originMeasurement string
)

func init() {
	value = flag.Float64("value", 10.07, "Adding value to converter")
	measurement = flag.String("measurement", "celsius or fahrenheit", "Adding conversor")
	flag.Parse()

	validArgs()

	originMeasurement = os.Args[lenArgs-1]
	originValues = os.Args[1 : lenArgs-1]
}

func validArgs() {
	lenArgs = len(os.Args)

	if lenArgs < 3 {
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

func calculateCelsius(originValue float64) float64 {
	return originValue*1.8 + 32
}

func calculateKilometre(originValue float64) float64 {
	return originValue / 1.60934
}

func destinyValue(originValue float64) (destinyValue float64) {
	if originMeasurement == "celsius" {
		destinyValue = calculateCelsius(originValue)
	} else {
		destinyValue = calculateKilometre(originValue)
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
