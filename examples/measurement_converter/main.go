package main

import (
	"flag"
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
	//	if len(os.Args) < 3 {
	//		fmt.Println("Use: converter <value> <measurement>")
	//		os.Exit(1)
	//	}
}
